package placetopay

//CollectRequest structure
type CollectRequest struct {
	Player     *Person         `json:"player"`
	Payment    *PaymentRequest `json:"payment"`
	Instrument *Instrument     `json:"instrument"`
}
