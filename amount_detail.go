package placetopay

//AmountDetail structure
type AmountDetail struct {
	Kind   string  `json:"kint,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	Base   float32 `json:"base,omitempty"`
}

type AmountBody struct {
	Currency string          `json:"currency,omitempty"`
	Total    float64         `json:"total,omitempty"`
	Taxes    []*TaxDetail    `json:"taxes,omitempty"`
	Details  []*AmountDetail `json:"details,omitempty"`
}
