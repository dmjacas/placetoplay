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
	AllResponse    string    `json:"allResponse" gorm:"size:5550"`
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

// PlacetoPayGetInformationLog structure
type PlacetoPayGetInformationLog struct {
	ID                int       `json:"id" gorm:"PRIMARY_KEY; AUTO_INCREMENT;size:11" `
	Active            bool      `json:"active"`
	RequestID         string    `json:"requestId"`
	AllResponse       string    `json:"allResponse" gorm:"size:5550"`
	Status            string    `json:"status"`
	Reason            string    `json:"reason"`
	Message           string    `json:"message"`
	InternalReference string    `json:"internalReference"`
	Authorization     string    `json:"authorization"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	DeletedAt         null.Time `json:"-"`
}

// PlacetoPayReversePaymemt structure
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
