package placetopay

import (
	"time"

	null "gopkg.in/guregu/null.v3"
)

// Purchase Model
type Purchase struct {
	ID         int       `json:"id"`
	Active     bool      `json:"active"`
	UserID     int       `json:"name"`
	Locale     string    `json:"locale"`
	Buyer      string    `json:"buyer"`
	Payment    string    `json:"payment"`
	Expiration string    `expiration:"expiration"`
	Response   string    `response:"response"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeletedAt  null.Time `json:"-"`
}

// NewPurchaseParams params passed to NewPlayer function
type NewPurchaseParams struct {
	UserID   int
	Locale   string
	Buyer    string
	Payment  string
	Type     string
	Response string
}

// NewPurchase create a new Purchase
func NewPurchase(params *NewPurchaseParams) *Purchase {
	return &Purchase{
		Active:   true,
		UserID:   params.UserID,
		Locale:   params.Locale,
		Response: params.Response,
		Buyer:    params.Buyer,
		Payment:  params.Payment,
		Type:     params.Type,
	}
}
