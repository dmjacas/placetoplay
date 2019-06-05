package placetopay

//Amount structure
type Amount struct {
	Currency string  `json:"currency"`
	Total    float64 `json:"total"`
	//Taxes    *TaxDetail    `json:"taxes"`
	// Details  *AmountDetail `json:"details"`
}
