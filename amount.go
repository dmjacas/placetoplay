package placetopay

//Amount structure
type Amount struct {
	Currency string          `json:"currency,omitempty"`
	Total    float64         `json:"total,omitempty"`
	Taxes    *[]TaxDetail    `json:"taxes,omitempty"`
	Details  *[]AmountDetail `json:"details,omitempty"`
}
