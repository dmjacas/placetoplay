package placetopay

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func init() {
	println("main package initialized")

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
func AuthRequest(login, Secret string) Auth {
	date := time.Now()
	nonce := RandStringBytesRmndr(16)
	seed := date.Format(time.RFC3339)
	tranKey := encodeBase64(sha1Encode(nonce + string(seed) + Secret))
	auth := Auth{
		Login:   login,
		Nonce:   nonce,
		TranKey: tranKey,
		Seed:    seed,
	}
	return auth
}

func CreateRequest(auth Auth, data RedirectRequest) RedirectResponse {

	ret := RedirectResponse{}
	fmt.Println("Geometrical shape properties")
	return ret
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
