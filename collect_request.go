package placetopay

//CollectRequest structure
type CollectRequest struct {
	Player     *Person         `json:"player"`
	Payment    *PaymentRequest `json:"payment"`
	Instrument *Instrument     `json:"instrument"`
}

// CollectBodyRequest
type CollectBodyRequest struct {
	Auth       *Auth        `json:"auth"`
	Payer      *Person      `json:"payer"`
	Payment    *PaymentBody `json:"paymemt"`
	Instrument *Instrument  `json:"instrument"`
}
