package placetopay

import (
	"time"

	null "gopkg.in/guregu/null.v3"
)

// PlacetoPayRequestLog Model
type PlacetoPayRequestLog struct {
	ID             int       `json:"id" gorm:"PRIMARY_KEY; AUTO_INCREMENT;size:11" `
	Active         bool      `json:"active"`
	Reference      string    `json:"reference"`
	AllResponse    string    `json:"allResponse" gorm:"size:2550"`
	Expiration     string    `json:"expiration" gorm:"size:2550"`
	IPAddress      string    `json:"ipadres"`
	ReturnURL      string    `json:"returnUrl" gorm:"size:550"`
	CancelURL      string    `json:"cancelUrl" gorm:"size:550"`
	SkipResult     bool      `json:"skipResult" `
	NoBuyerFill    bool      `json:"noBuyerFill"`
	CaptureAddress bool      `json:"captureAddress"`
	PaymentMethod  bool      `json:"paymentMethod"`
	Fields         string    `json:"fields" gorm:"size:2550"`
	RequestID      string    `json:"requestId"`
	ProcessURL     string    `json:"processUrl" gorm:"size:250"`
	Message        string    `json:"message" gorm:"size:250"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	DeletedAt      null.Time `json:"-"`
}

//PlacetoPayGetInformationLog structure
type PlacetoPayGetInformationLog struct {
	ID                int       `json:"id" gorm:"PRIMARY_KEY; AUTO_INCREMENT;size:11" `
	Active            bool      `json:"active"`
	RequestID         string    `json:"requestId"`
	AllResponse       string    `json:"allResponse" gorm:"size:5550"`
	Status            string    `json:"status"`
	Reason            string    `json:"reason"`
	Message           string    `json:"message"`
	InternalReference float64   `json:"internalReference"`
	Authorization     string    `json:"authorization"`
	PaymentMethod     string    `json:"paymentMethod" gorm:"size:250"`
	PaymentMethodName string    `json:"paymentMethodName" gorm:"size:250"`
	IssuerName        string    `json:"issuerName" gorm:"size:250"`
	Reference         string    `json:"reference" gorm:"size:250"`
	Receipt           string    `json:"receipt" gorm:"size:250"`
	Code              string    `json:"code" gorm:"size:250"`
	Installments      string    `json:"installments" gorm:"size:250"`
	GroupCode         string    `json:"groupCode" gorm:"size:250"`
	CodeLegend        string    `json:"codeLegend" gorm:"size:250"`
	LastDigits        string    `json:"lastDigits"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	DeletedAt         null.Time `json:"-"`
}

//PlacetoPayReversePaymemt structure
type PlacetoPayReversePaymemt struct {
	ID        int       `json:"id" gorm:"PRIMARY_KEY; AUTO_INCREMENT;size:11" `
	Active    bool      `json:"active"`
	Status    string    `json:"status"`
	Reason    string    `json:"reason"`
	Message   string    `json:"message"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt null.Time `json:"-"`
}
