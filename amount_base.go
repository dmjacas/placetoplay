package placetopay

//AmountBase structure
type AmountBase struct {
	Currency string  `json:"currency"`
	Total    float64 `json:"total"`
}
