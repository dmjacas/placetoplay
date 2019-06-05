package placetopay

// PaymentRequest structure
type PaymentRequest struct {
	Reference    string           `json:"reference"`
	Description  string           `json:"description"`
	Amount       *Amount          `json:"amount"`
	AllowPartial bool             `json:"allowPartial"`
	Shipping     *Person          `json:"shipping"`
	Items        []*Item          `json:"items"`
	Fields       []*NameValuePair `json:"fields"`
	Recurring    *Recurring       `json:"recurring"`
	Subcribe     bool             `json:"subcribe"`
}
