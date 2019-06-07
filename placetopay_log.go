package placetopay

import (
	"time"

	null "gopkg.in/guregu/null.v3"
)

// PlacetoPayRequestLog Model
type PlacetoPayRequestLog struct {
	ID             int       `json:"id"`
	Active         bool      `json:"active"`
	Reference      string    `json:"reference"`
	Buyer          string    `json:"buyer" gorm:"size:2550"`
	Payment        string    `json:"payment" gorm:"size:2550"`
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
	Response       string    `json:"response" gorm:"size:2550"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	DeletedAt      null.Time `json:"-"`
}

// NewPlacetoPayRequestLogParams params passed to NewPlayer function
type NewPlacetoPayRequestLogParams struct {
	Buyer          string
	Reference      string
	Payment        string
	Expiration     string
	IPAddress      string
	ReturnURL      string
	CancelURL      string
	SkipResult     bool
	NoBuyerFill    bool
	CaptureAddress bool
	PaymentMethod  bool
	Fields         string
	RequestID      string
	ProcessURL     string
	Message        string
	Response       string
}

// NewPlacetoPayRequestLog create a new Purchase
func NewPlacetoPayRequestLog(params *NewPlacetoPayRequestLogParams) *PlacetoPayRequestLog {
	return &PlacetoPayRequestLog{
		Active:         true,
		Payment:        params.Payment,
		Reference:      params.Reference,
		Expiration:     params.Expiration,
		IPAddress:      params.IPAddress,
		ReturnURL:      params.ReturnURL,
		CancelURL:      params.CancelURL,
		SkipResult:     params.SkipResult,
		NoBuyerFill:    params.NoBuyerFill,
		CaptureAddress: params.CaptureAddress,
		PaymentMethod:  params.PaymentMethod,
		Fields:         params.Fields,
		RequestID:      params.RequestID,
		ProcessURL:     params.ProcessURL,
		Message:        params.Message,
		Response:       params.Response,
	}
}

// PlacetoPayGetInformationLog structure
type PlacetoPayGetInformationLog struct {
	ID                int       `json:"id"`
	Active            bool      `json:"active"`
	RequestID         string    `json:"requestId"`
	AllStatus         string    `json:"allStatus" gorm:"size:2550"`
	AllRequest        string    `json:"allRequest" gorm:"size:2550"`
	AllPayment        string    `json:"AllPayment" gorm:"size:2550"`
	AllSubscription   string    `json:"AllSubscription" gorm:"size:2550"`
	Status            string    `json:"status"`
	Reason            string    `json:"reason"`
	Message           string    `json:"message"`
	InternalReference string    `json:"internalReference"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	DeletedAt         null.Time `json:"-"`
}

// NewPlacetoPayGetInformationLogParams params passed to NewPlayer function
type NewPlacetoPayGetInformationLogParams struct {
	RequestID         string
	AllStatus         string
	AllRequest        string
	AllPayment        string
	AllSubscription   string
	Status            string
	Reason            string
	Message           string
	InternalReference string
}

// NewPlacetoPayGetInformationLog create a new Purchase
func NewPlacetoPayGetInformationLog(params *NewPlacetoPayGetInformationLogParams) *PlacetoPayGetInformationLog {
	return &PlacetoPayGetInformationLog{
		Active:          true,
		RequestID:       params.RequestID,
		AllStatus:       params.AllStatus,
		AllRequest:      params.AllRequest,
		AllPayment:      params.AllPayment,
		AllSubscription: params.AllSubscription,
		Status:          params.Status,
		Reason:          params.Reason,
		Message:         params.Message,

		InternalReference: params.InternalReference,
	}
}

type PlacetoPayReversePaymemt struct {
	ID                int       `json:"id"`
	Active            bool      `json:"active"`
	InternalReference string    `json:"internalReference"`
	Reference         string    `json:"reference" `
	AllReference      string    `json:"allReference" gorm:"size:2550"`
	Status            string    `json:"status"`
	Reason            string    `json:"reason"`
	Message           string    `json:"message"`
	AllStatus         string    `json:"allStatus" gorm:"size:2550"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	DeletedAt         null.Time `json:"-"`
}
