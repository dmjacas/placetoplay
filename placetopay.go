package placetopay

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"time"
)

var URLPayment string
var URLReturn string
var Login string
var Secret string
var DBAConection string

// Config config payment library
func Config(Payment, Return, secret, login string) {
	URLReturn = Return
	URLPayment = Payment
	Login = login
	Secret = secret

}
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
func AuthRequest(login, Secret string) *Auth {
	date := time.Now()
	nonce := RandStringBytesRmndr(16)
	nonceEncode := encodeBase64(sha1Encode(nonce))
	fmt.Println(nonce)
	fmt.Println(nonceEncode)
	seed := date.Format(time.RFC3339)
	tranKey := encodeBase64(sha1Encode(nonce + string(seed) + Secret))
	auth := &Auth{
		Login:   login,
		Nonce:   nonceEncode,
		TranKey: tranKey,
		Seed:    seed,
	}
	return auth
}

// CreateRequest Create paymente request
func CreateRequest(data *RedirectRequest) *RedirectRequest {
	auth := AuthRequest(Login, Secret)
	data.Auth = auth

	jsonRequest, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error json")
	}
	fmt.Println(URLPayment)

	ret, err := http.Post(URLPayment, "application/json", bytes.NewBuffer(jsonRequest))
	if err != nil {
		fmt.Println("error 1")
	}
	dat, _ := ioutil.ReadAll(ret.Body)
	var retorno = RedirectResponse{}
	if err = json.Unmarshal(dat, &retorno); err != nil {
		fmt.Println("error 2")
	}
	fmt.Println(ret.Body)
	return data
}
func main() {
	fmt.Println("Geometrical shape properties")

}

//GetIPAddress IP address from server
func GetIPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
