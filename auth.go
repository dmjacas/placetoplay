package placetopay

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

// Auth structure
type Auth struct {
	Login      string `json:"login"`
	Nonce      string `json:"nonce"`
	Seed       string `json:"seed"`
	TranKey    string `json:"tranKey"`
	Expiration string `json:"expiration"`
}

/*
// AuthBody structure
type AuthBody struct {
	ConnectionType string `json:"connectionType"`
	Endpoint       string `json:"endpoint"`
	Login          string `json:"login"`
	TranKey        string `json:"tranKey"`
}*/

type DBConfig struct {
	Dialect  string `default:"mysql"`
	Username string
	Password string
	Name     string
	Charset  string
}

// Connect handles the connection to the database and returns it
func Connect(config *DBConfig) (*gorm.DB, error) {

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

var URLPayment string
var URLReturn string
var Login string
var Secret string
var DBAConection string
var DB *gorm.DB

// Config config payment library
func Config(Payment, Return, secret, login, dbCharset, dbDialect, dbName, dbPassword, dbUsername string) {
	URLReturn = Return
	URLPayment = Payment
	Login = login
	Secret = secret
	db := &DBConfig{
		Dialect:  dbDialect,
		Username: dbUsername,
		Password: dbPassword,
		Name:     dbName,
		Charset:  dbCharset,
	}
	DB, _ = Connect(db)
	fmt.Println("config db " + dbName)

}
