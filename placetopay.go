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
	"strconv"
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
// login PlacetoPay login
// secret PlacetoPay secret password
// dbCharset db Charset
// dbDialect db Dialect
// dbName dn name
// dbUsername db username
// dbPassword db password
// Expiration time in minutes that the payment request lasts

// Config pay db connectiong
func Config(urlPayment, login, secret, dbCharset, dbDialect, dbName, dbUsername, dbHost, dbPort, dbPassword string, Expiration int) {
	P2PURLPayment = urlPayment
	P2PLogin = login
	P2PSecret = secret
	P2PExpiration = Expiration
	db := &dBConfig{
		Dialect:  dbDialect,
		Username: dbUsername,
		Password: dbPassword,
		Host:     dbHost,
		Name:     dbName,
		Port:     dbPort,
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
		P2PDB.AutoMigrate(&PlacetoPayRequestLog{}, &PlacetoPayGetInformationLog{}, &PlacetoPayReversePaymemt{})
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
	//stringResponse, err := json.Marshal(placeToPayResponse)
	if err != nil {
		return nil, errors.New("error to convert to string Response Data")
	}

	stringFields, err := json.Marshal(data.Buyer)
	if err != nil {
		return nil, errors.New("error to JSON encode the body request")
	}

	stringResponse, err := json.Marshal(placeToPayResponse)
	if err != nil {
		return nil, errors.New("error to JSON encode the body request")
	}
	tx := P2PDB.Begin()
	// save the log of the payment request
	requestLog := PlacetoPayRequestLog{
		Active:         true,
		Reference:      data.Payment.Reference,
		AllResponse:    string(stringResponse),
		Expiration:     data.Expiration,
		IPAddress:      data.IPAddres,
		ReturnURL:      data.ReturnURL,
		CancelURL:      data.ReturnURL,
		SkipResult:     data.SkipResult,
		NoBuyerFill:    data.NoBuyerFill,
		CaptureAddress: false,
		PaymentMethod:  false,
		Fields:         string(stringFields),
		RequestID:      strconv.Itoa(placeToPayResponse.RequestID),
		ProcessURL:     placeToPayResponse.ProcessURL,
		Message:        placeToPayResponse.Status.Message,
	}
	if result := tx.Create(&requestLog); result.Error != nil {
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
func GetRequestInformation(requestID string) (*RedirectInformation, error) {
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

	var placeToPayResponse RedirectInformation

	if err = json.Unmarshal(dataResp, &placeToPayResponse); err != nil {
		return nil, errors.New("error in convert response to RedirectResponse")
	}

	stringResponse, err := json.Marshal(placeToPayResponse)
	if err != nil {
		return nil, errors.New("error to JSON encode the body request")
	}
	InternalReference := ""
	Authorization := ""
	if placeToPayResponse.Payment[0] != nil {
		InternalReference = placeToPayResponse.Payment[0].Receipt
		Authorization = placeToPayResponse.Payment[0].Authorization
	}
	tx := P2PDB.Begin()
	// save the log of the payment request
	requestLog := &PlacetoPayGetInformationLog{
		Active:            true,
		RequestID:         requestID,
		AllResponse:       string(stringResponse),
		Status:            placeToPayResponse.Status.Status,
		Reason:            placeToPayResponse.Status.Reason,
		Message:           placeToPayResponse.Status.Message,
		InternalReference: InternalReference,
		Authorization:     Authorization,
	}

	if result := tx.Create(&requestLog); result.Error != nil {
		tx.Rollback()
		return nil, errors.New("error in saving the data")
	}
	if result := tx.Commit(); result.Error != nil {
		tx.Rollback()
		return nil, errors.New("error in saving the data")
	}
	return &placeToPayResponse, nil

}

// ReversePayment reverte payment request
// requestID request id
func ReversePayment(internalReference string) (*ReverseResponse, error) {
	// Get auth object
	auth, _ := authRequest()
	// Generate body request
	bodyRequest := &ReverseBodyRequest{
		Auth:              auth,
		InternalReference: internalReference,
	}
	// Encode JSON  body request
	jsonRequest, err := json.Marshal(bodyRequest)
	if err != nil {
		return nil, errors.New("error to JSON encode the body request")
	}
	// call the P2P api
	response, err := http.Post(P2PURLPayment+"api/reverse", "application/json", bytes.NewBuffer(jsonRequest))
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

	tx := P2PDB.Begin()
	// save the log of the payment request
	requestLog := &PlacetoPayReversePaymemt{
		Active:  true,
		Status:  placeToPayResponse.Status.Status,
		Reason:  placeToPayResponse.Status.Reason,
		Message: placeToPayResponse.Status.Message,
		Date:    placeToPayResponse.Status.Date,
	}

	if result := tx.Create(&requestLog); result.Error != nil {
		tx.Rollback()
		return nil, errors.New("error in saving the data")
	}
	if result := tx.Commit(); result.Error != nil {
		tx.Rollback()
		return nil, errors.New("error in saving the data")
	}
	return &placeToPayResponse, nil
}

//CollectPayment implementar v1.2
/*func CollectPayment(colection *CollectBodyRequest) RedirectInformation {
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
}*/

/*
*	Complements functions
 */

// Connect handles the connection to the database and returns it
func Connect(config *dBConfig) (*gorm.DB, error) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.Charset)
	db, err := gorm.Open(config.Dialect, dbURI)
	if err != nil {
		log.Fatalln("db connect", err)
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
	Host     string `default:"localhost"`
	Port     string `default:"3306"`
	Dialect  string `default:"mysql"`
	Username string
	Password string
	Name     string
	Charset  string `default:"utf8"`
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

//Person structure
type Person struct {
	DocumenType string  `json:"documentType,omitempty"`
	Document    string  `json:"document,omitempty"`
	Name        string  `json:"name,omitempty"`
	Surname     string  `json:"surname,omitempty"`
	Company     string  `json:"company,omitempty"`
	Email       string  `json:"email,omitempty"`
	Addres      *Addres `json:"addres,omitempty"`
	Mobile      string  `json:"mobile,omitempty"`
}

// Addres structure
type Addres struct {
	Street     string `json:"street,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
	Country    string `json:"country,omitempty"`
	Phone      string `json:"phone,omitempty"`
}

// RedirectRequest structure
type RedirectRequest struct {
	Auth          *Auth                `json:"auth,omitempty"`
	Locale        string               `json:"locale,omitempty" `
	Payer         *Person              `json:"payer,omitempty"`
	Buyer         *Person              `json:"buyer,omitempty"`
	Payment       *PaymentRequest      `json:"payment,omitempty"`
	Subscription  *SubscriptionRequest `json:"subscription,omitempty"`
	Fields        []*NameValuePair     `json:"fields,omitempty"`
	PaymentMethod string               `json:"paymentMethod,omitempty"`
	Expiration    string               `json:"expiration,omitempty"`
	ReturnURL     string               `json:"returnUrl,omitempty"`
	IPAddres      string               `json:"ipAddress,omitempty"`
	UserAgent     string               `json:"userAgent,omitempty"`
	SkipResult    bool                 `json:"skipResult,omitempty"`
	NoBuyerFill   bool                 `json:"noBuyerFill,omitempty"`
}

// RedirectResponse structure
type RedirectResponse struct {
	Status     *Status `json:"status,"`
	RequestID  int     `json:"requestId"`
	ProcessURL string  `json:"processUrl"`
}

// RedirectInformation structure
type RedirectInformation struct {
	Status       *Status                   `json:"status"`
	Request      *RedirectRequest          `json:"request"`
	Payment      []*BodyPaymentInformation `json:"payment"`
	Subscription *SubscriptionResponse     `json:"subscription"`
}

//BodyPaymentInformation structure
type BodyPaymentInformation struct {
	Status            *Status `json:"status,omitempty"`
	InternalReference int64   `json:"internalReference,omitempty"`
	PaymentMethod     string  `json:"paymentMethod,omitempty"`
	PaymentMethodName string  `json:"paymentMethodName,omitempty"`
	IssuerName        string  `json:"issuerName,omitempty"`
	Reference         string  `json:"reference,omitempty"`
	Authorization     string  `json:"authorization,omitempty"`
	Receipt           string  `json:"receipt,omitempty"`
}

//AmountBase structure
type AmountBase struct {
	Currency string  `json:"currency,omitempty"`
	Total    float64 `json:"total,omitempty"`
}

//AmountConversion structure
type AmountConversion struct {
	From   *AmountBase `json:"from"`
	To     *AmountBase `json:"to"`
	Factor float64     `json:"factor"`
}

//Transaction structure
type Transaction struct {
	Status            *Status           `json:"status,omitempty"`
	InternalReference string            `json:"internalReference,omitempty"`
	Reference         string            `json:"reference,omitempty"`
	PaymentMethod     string            `json:"paymentMethod,omitempty"`
	PaymentMethodName string            `json:"paymentMethodName,omitempty"`
	IssuerName        string            `json:"issuerName,omitempty"`
	Amount            *AmountConversion `json:"amount,omitempty"`
	Receipt           string            `json:"receipt,omitempty"`
	Frachise          string            `json:"frachise,omitempty"`
	Refunded          bool              `json:"refunded,omitempty"`
	Authorization     string            `json:"authorization,omitempty"`
	ProcessorFields   *NameValuePair    `json:"processorFields,omitempty"`
}

// ReverseResponse structure
type ReverseResponse struct {
	Status  *Status      `json:"status,omitempty"`
	Payment *Transaction `json:"payment,omitempty"`
}

// ReverseBodyRequest structure
type ReverseBodyRequest struct {
	Auth              *Auth  `json:"auth,omitempty"`
	InternalReference string `json:"internalReference,omitempty"`
}

//SubscriptionRequest structure
type SubscriptionRequest struct {
	Reference   string         `json:"reference,omitempty"`
	Description string         `json:"description,omitempty"`
	Fields      *NameValuePair `json:"fields,omitempty"`
}

//SubscriptionResponse structure
type SubscriptionResponse struct {
	Status    *Status        `json:"reference,omitempty"`
	Type      string         `json:"type,omitempty"`
	Intrument *NameValuePair `json:"fields,omitempty"`
}

// SimpleToken structure
type SimpleToken struct {
	Token        string `json:"token,omitempty"`
	Subtoken     string `json:"subtoken,omitempty"`
	Installments int    `json:"installments,omitempty"`
	Cvv          string `json:"cvv,omitempty"`
}

//Recurring structure
type Recurring struct {
	Periodicity     string    `json:"periodicity,omitempty"`
	Interval        int       `json:"interval,omitempty"`
	NextPayment     time.Time `json:"nextPayment,omitempty"`
	MaxPeriods      int       `json:"maxPeriods,omitempty"`
	DueDate         time.Time `json:"dueDate,omitempty"`
	NotificationURL string    `json:"notificationUrl,omitempty"`
}

//Item structure
type Item struct {
	Sku      string  `json:"street,omitempty"`
	Name     string  `json:"name,omitempty"`
	Category string  `json:"category,omitempty"`
	Qty      string  `json:"qty,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Tax      float64 `json:"tax,omitempty"`
}

//DocumentType structure
type DocumentType struct {
	Country      string `json:"country,omitempty"`
	Code         string `json:"code,omitempty"`
	DocumentType string `json:"documentType,omitempty"`
}

// PaymentRequest structure
type PaymentRequest struct {
	Reference    string           `json:"reference,omitempty"`
	Description  string           `json:"description,omitempty"`
	Amount       *Amount          `json:"amount,omitempty"`
	AllowPartial bool             `json:"allowPartial,omitempty"`
	Shipping     *Person          `json:"shipping,omitempty"`
	Items        []*Item          `json:"items,omitempty"`
	Fields       []*NameValuePair `json:"fields,omitempty"`
	Recurring    *Recurring       `json:"recurring,omitempty"`
	Subcribe     bool             `json:"subcribe,omitempty"`
}

// Instrument structure
type Instrument struct {
	Token string `json:"token"`
}

//NameValuePair structure
type NameValuePair struct {
	Keyword   string `json:"keyword,omitempty"`
	Value     string `json:"value,omitempty"`
	DisplayOn string `json:"displayOn,omitempty"`
}

//AmountDetail structure
type AmountDetail struct {
	Kind   string  `json:"kint,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	Base   float32 `json:"base,omitempty"`
}

//Amount structure
type Amount struct {
	Currency string          `json:"currency,omitempty"`
	Total    float64         `json:"total,omitempty"`
	Taxes    []*TaxDetail    `json:"taxes,omitempty"`
	Details  []*AmountDetail `json:"details,omitempty"`
}

//CollectRequest structure
type CollectRequest struct {
	Player     *Person         `json:"player"`
	Payment    *PaymentRequest `json:"payment"`
	Instrument *Instrument     `json:"instrument"`
}

// CollectBodyRequest structure
type CollectBodyRequest struct {
	Auth       *Auth        `json:"auth"`
	Payer      *Person      `json:"payer"`
	Payment    *PaymentBody `json:"paymemt"`
	Instrument *Instrument  `json:"instrument"`
}

//PaymentBody structure
type PaymentBody struct {
	Reference   string  `json:"reference"`
	Description string  `json:"description"`
	Amount      *Amount `json:"amount" binding:"required,dive"`
}
