package placetopay

type PlacetoPayBodyRequest struct {
	Auth          *AuthBody    `json:"auth"`
	Buyer         *BuyerBody   `json:"buyer"`
	Payment       *PaymentBody `json:"payment"`
	Expiration    string       `json:"expiration"`
	IPAddress     string       `json:"ipAddress"`
	ReturnURL     string       `json:"returnUrl"`
	UseAgent      string       `json:"userAgent"`
	PaymentMethod string       `json:"paymentMethod"`
}
type PlacetoPayStatusBodyRequest struct {
	Auth *AuthBody `json:"auth"`
}
type AuthBody struct {
	Login   string `json:"login"`
	Nonce   string `json:"nonce"`
	Seed    string `json:"seed"`
	TranKey string `json:"tranKey"`
}

type BuyerBody struct {
	Name         string   `json:"name"`
	Surname      string   `json:"surname"`
	Email        string   `json:"email"`
	Document     string   `json:"document"`
	DocumentType string   `json:"documentType"`
	Mobile       string   `json:"mobile"`
	Phone        string   `json:"phone"`
	Address      *Address `json:"address"`
}

type Address struct {
	Phone  string `json:"phone"`
	Street string `json:"street"`
}
type PaymentBody struct {
	Reference   string      `json:"reference"`
	Description string      `json:"description"`
	Amount      *AmountBody `json:"amount" binding:"required,dive"`
}

type TaxesBody struct {
	Kind   string  `json:"kind"`
	Amount float64 `json:"amount"`
	Base   float64 `json:"base"`
}

type AmountDetailBody struct {
	Kind   string  `json:"kind"`
	Amount float64 `json:"amount"`
}

type AmountBody struct {
	Currency string              `json:"currency"`
	Total    float64             `json:"total"`
	Taxes    []*TaxesBody        `json:"taxes"`
	Details  []*AmountDetailBody `json:"details"`
}

type PlacetoPayBodyResponse struct {
	Status     *StatusBody `json:"status"`
	RequestID  int         `json:"requestId"`
	ProcessURL string      `json:"processUrl"`
}
type PlacetoPayStatusBodyResponse struct {
	Status    *StatusBody `json:"status"`
	RequestID int         `json:"requestId"`
	Request   interface{} `json:"request"`
	Payment   interface{} `json:"payment"`
}
type StatusBody struct {
	Status  string `json:"status" binding:"required"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
	Date    string `json:"date" binding:"required"`
}
type NotificationBody struct {
	Status    *StatusBody `json:"status"`
	RequestID int64       `json:"requestId" binding:"required"`
	Reference string      `json:"reference" binding:"required"`
	Signature string      `json:"signature" binding:"required"`
}

type Configuration struct {
	RequestURL string `json:"request_url" binding:"required,url"`
	Currency   string `json:"currency" binding:"required"`
	Login      string `json:"login" binding:"required"`
	Secret     string `json:"secret" binding:"required"`
	RetunrURL  string `json:"return_url" binding:"required,url"`
}
