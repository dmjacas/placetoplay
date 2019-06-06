package placetopay

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

// P2PURLPayment PlacetoPay api url
var P2PURLPayment string

// P2PLogin PlacetoPlay login
var P2PLogin string

//P2PSecret PlacetoPlay scret string
var P2PSecret string

//P2PExpiration PlacetoPlay transaction expiration time
var P2PExpiration int

// P2PDB db object
var P2PDB *gorm.DB
var db *dBConfig

// Config configure payment library
// urlPayment PlacetoPay api url
// secret PlacetoPay secret password
// login PlacetoPay login
// dbCharset db Charset
// dbDialect db Dialect
// dbName dn name
// dbPassword db password
// dbUsername db username
func Config(urlPayment, secret, login, dbCharset, dbDialect, dbName, dbPassword, dbUsername string) {
	P2PURLPayment = urlPayment
	P2PLogin = login
	P2PSecret = secret
	P2PExpiration = 10
	db := &dBConfig{
		Dialect:  dbDialect,
		Username: dbUsername,
		Password: dbPassword,
		Name:     dbName,
		Charset:  dbCharset,
	}
	P2PDB, _ = Connect(db)
	migration()
}

// migration  create table if not exist
func migration() {
	pingErr := P2PDB.DB().Ping()
	if pingErr != nil {
		fmt.Println(pingErr)
	} else {
		P2PDB.AutoMigrate(&Purchase{})
	}
}

// letterBytes all letters
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 6 bits to represent a letter index
const letterIdxBits = 6

// All 1-bits, as many as letterIdxBits
const letterIdxMask = 1<<letterIdxBits - 1

// CreateRequest Create paymente request
func CreateRequest(data *RedirectRequest) (*RedirectResponse, error) {
	auth, expiration := authRequest()
	data.Auth = auth
	data.Expiration = expiration
	jsonRequest, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("error to JSON encode the body request")
	}
	response, err := http.Post(P2PURLPayment+"api/session", "application/json", bytes.NewBuffer(jsonRequest))
	if err != nil {
		return nil, errors.New("error in the http call")
	}
	dataResp, err := ioutil.ReadAll(response.Body)
	var placeToPayResponse RedirectResponse
	if err = json.Unmarshal(dataResp, &placeToPayResponse); err != nil {
		return nil, errors.New("error in the return values of the http call")
	}
	stringBuyer, err := json.Marshal(data.Buyer)
	if err != nil {
		return nil, errors.New("error to convert to string BuyerData")
	}
	stringPayment, err := json.Marshal(data.Payment)
	if err != nil {
		return nil, errors.New("error to convert to string PaymentData")
	}
	stringResponse, err := json.Marshal(placeToPayResponse)
	if err != nil {
		return nil, errors.New("error to convert to string Response Data")
	}

	tx := P2PDB.Begin()

	// save the log of the payment request
	purchase := NewPurchase(&NewPurchaseParams{
		Buyer:    string(stringBuyer),
		Locale:   "es_CO",
		Payment:  string(stringPayment),
		Response: string(stringResponse),
		Type:     "",
		UserID:   3,
	})
	if result := tx.Create(&purchase); result.Error != nil {
		tx.Rollback()
		return nil, errors.New("error in saving the data")
	}
	if result := tx.Commit(); result.Error != nil {
		tx.Rollback()
		return nil, errors.New("error in saving the data")
	}
	return &placeToPayResponse, nil
}

// GetRequestInformation get the status information of the request
// requestID request id
func GetRequestInformation(requestID string) (*RedirectResponse, error) {
	auth, _ := authRequest()
	bodyRequest := &StatusBodyRequest{
		Auth: auth,
	}
	jsonRequest, err := json.Marshal(bodyRequest)
	if err != nil {
		return nil, errors.New("error to JSON encode the body request")
	}
	response, err := http.Post(P2PURLPayment+"api/session/"+requestID, "application/json", bytes.NewBuffer(jsonRequest))
	if err != nil {
		return nil, errors.New("error in the http call")
	}
	dataResp, _ := ioutil.ReadAll(response.Body)

	var placeToPayResponse RedirectResponse

	if err = json.Unmarshal(dataResp, &placeToPayResponse); err != nil {
		return nil, errors.New("error in convert response to RedirectResponse")
	}
	return &placeToPayResponse, nil

}

// ReversePaymemt reverte payment request
// requestID request id
func ReversePaymemt(requestID string) (*ReverseResponse, error) {
	// Get auth object
	auth, _ := authRequest()
	// Generate body request
	bodyRequest := &ReverseBodyRequest{
		Auth:              auth,
		InternalReference: requestID,
	}
	// Encode JSON  body request
	jsonRequest, err := json.Marshal(bodyRequest)
	if err != nil {
		return nil, errors.New("error to JSON encode the body request")
	}
	// call the P2P api
	response, err := http.Post(P2PURLPayment+"api/reverse/", "application/json", bytes.NewBuffer(jsonRequest))
	if err != nil {
		return nil, errors.New("error in the http call")
	}
	// Get response
	dataResp, _ := ioutil.ReadAll(response.Body)

	var placeToPayResponse ReverseResponse
	// Convert response to ReverseResponse object
	if err = json.Unmarshal(dataResp, &placeToPayResponse); err != nil {
		return nil, errors.New("error in convert response to ReverseResponse")
	}
	return &placeToPayResponse, nil
}

//CollectPayment (falta)
func CollectPayment(colection *CollectBodyRequest) RedirectInformation {
	auth, _ := authRequest()

	colection.Auth = auth

	jsonRequest, err := json.Marshal(colection)
	if err != nil {
		fmt.Println("error json")
	}
	response, err := http.Post(P2PURLPayment+"api/collect/", "application/json", bytes.NewBuffer(jsonRequest))
	if err != nil {
		fmt.Println("error http")
		// return
	}

	dataResp, _ := ioutil.ReadAll(response.Body)

	var placeToPayResponse RedirectInformation

	if err = json.Unmarshal(dataResp, &placeToPayResponse); err != nil {
		fmt.Println("error http")
		//	return
	}
	return placeToPayResponse
}

/*
*	Complements functions
 */

// Connect handles the connection to the database and returns it
func Connect(config *dBConfig) (*gorm.DB, error) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.Username,
		config.Password,
		config.Name,
		config.Charset)

	db, err := gorm.Open(config.Dialect, dbURI)
	if err != nil {
		log.Fatalln("aqui", err)
	}

	return db, nil
}

// createRandString create a random string
// lengt integer value of the number of characters in the string
func createRandString(lengt int) string {

	randString := make([]byte, lengt)
	for i := range randString {
		randString[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(randString)
}

// sha1Encode encode string  to  sh1
// str string to encode
func sha1Encode(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return string(bs)
}

// encodeBase64 encode string  to base64
// str string to encode

func encodeBase64(msg string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	return encoded
}

// authRequest create authentication request
func authRequest() (*Auth, string) {
	// get current time
	date := time.Now()
	// generate random string to 16 character
	nonce := createRandString(16)
	// encode (nonce) to base64
	nonceEncode := encodeBase64(nonce)
	// format current time to Date ISO 8601
	seed := date.Format(time.RFC3339)
	// tranKey
	//tranKey value generate with concatenate nonce + seed + Secret encoded to Base64
	tranKey := encodeBase64(sha1Encode(nonce + string(seed) + P2PSecret))
	// expiration
	expiration := date.Add(time.Minute * time.Duration(P2PExpiration)).Format(time.RFC3339)
	// auth initialize auth structure
	auth := &Auth{
		Login:   P2PLogin,
		Nonce:   nonceEncode,
		TranKey: tranKey,
		Seed:    seed,
	}
	// return auth structure and expiration time
	return auth, expiration
}

/**
* Estructures
 */

// Auth structure
type Auth struct {
	// Login PlactoPay login
	Login string `json:"login"`
	//Nonce encode Base64 random string
	Nonce string `json:"nonce"`
	//current date with format ISO 8601
	Seed string `json:"seed"`
	//TranKey  Base64(SHA-1(nonce + seed + tranKey))
	TranKey string `json:"tranKey"`
}

// dBConfig database config structure
type dBConfig struct {
	Dialect  string `default:"mysql"`
	Username string
	Password string
	Name     string
	Charset  string
}

// StatusBodyRequest body request structure
type StatusBodyRequest struct {
	Auth *Auth `json:"auth,omitempty"`
}

//TaxDetail structure
type TaxDetail struct {
	//
	Kind   string  `json:"kint,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	Base   float64 `json:"base,omitempty"`
}

// Status structure
type Status struct {
	Status  string    `json:"status,omitempty"`
	Reason  string    `json:"reason,omitempty"`
	Message string    `json:"message,omitempty"`
	Date    time.Time `json:"date,omitempty"`
}
