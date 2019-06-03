package placetopay

//Transaction structure
type Transaction struct {
	Status            *Status           `json:"status"`
	InternalReference int               `json:"internalReference"`
	Reference         string            `json:"reference"`
	PaymentMethod     string            `json:"paymentMethod"`
	PaymentMethodName string            `json:"paymentMethodName"`
	IssuerName        string            `json:"issuerName"`
	Amount            *AmountConversion `json:"amount"`
	Receipt           string            `json:"receipt"`
	Frachise          string            `json:"frachise"`
	Refunded          bool              `json:"refunded"`
	Authorization     string            `json:"authorization"`
	ProcessorFields   *NameValuePair    `json:"processorFields"`
}
