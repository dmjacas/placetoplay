package placetopay

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	const ReverseOrderStatus = "reversed"

}

// RandStringBytesRmndr create a random string
func RandStringBytesRmndr(n int) string {
	const (
		letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 52 possibilities
		letterIdxBits = 6
	)
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func sha1Encode(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return string(bs)
}

func encodeBase64(msg string) string {

	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	return encoded
}

func decodeBase64(str string) string {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("decode error:", err)
		return "error"
	}
	return string(decoded)
}

// AuthRequest create autentication request
func AuthRequest(login, Secret string) (*Auth, string) {
	date := time.Now()
	nonce := RandStringBytesRmndr(16)
	nonceEncode := encodeBase64(nonce)
	fmt.Println(nonce)
	fmt.Println(nonceEncode)
	seed := date.Format(time.RFC3339)
	fmt.Println(nonce + string(seed) + Secret)
	tranKey := encodeBase64(sha1Encode(nonce + string(seed) + Secret))
	expiration := date.Add(time.Minute * time.Duration(10)).Format(time.RFC3339)
	auth := &Auth{
		Login:   login,
		Nonce:   nonceEncode,
		TranKey: tranKey,
		Seed:    seed,
	}
	return auth, expiration
}

// CreateRequest Create paymente request
func CreateRequest(data *RedirectRequest) RedirectResponse {

	auth, expiration := AuthRequest(Login, Secret)
	data.Auth = auth
	data.Expiration = expiration

	jsonRequest, err := json.Marshal(data)

	if err != nil {
		fmt.Println("error json")
	}
	response, err := http.Post(URLPayment+"api/session", "application/json", bytes.NewBuffer(jsonRequest))
	if err != nil {
		fmt.Println("error http")
		// return
	}

	dataResp, _ := ioutil.ReadAll(response.Body)

	var placeToPayResponse RedirectResponse

	if err = json.Unmarshal(dataResp, &placeToPayResponse); err != nil {
		fmt.Println("error http")
		//	return
	}

	stringBuyer, err := json.Marshal(data.Buyer)

	if err != nil {
		fmt.Printf("Error: %s", err)
		//	return
	}
	stringPayment, err := json.Marshal(data.Payment)
	if err != nil {
		fmt.Printf("Error: %s", err)
		// return
	}

	stringResponse, err := json.Marshal(placeToPayResponse)
	if err != nil {
		fmt.Printf("Error: %s", err)
		// return
	}

	tx := DB.Begin()
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
		//c.Error(result.Error)
		// return
	}

	if result := tx.Commit(); result.Error != nil {
		tx.Rollback()
		// c.Error(result.Error)
		// return
	}
	return placeToPayResponse
}

// GetRequestInformation optiene la informacion del pago
func GetRequestInformation() {

}

func main() {
	fmt.Println("Geometrical shape properties")

}
