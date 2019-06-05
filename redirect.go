package placetopay

import "gopkg.in/guregu/null.v3"

// RedirectRequest structure
type RedirectRequest struct {
	Auth          *Auth                `json:"auth"`
	Locale        string               `json:"locale" `
	Payer         *Person              `json:"payer"`
	Buyer         *Person              `json:"buyer"`
	Payment       *PaymentRequest      `json:"payment"`
	Subscription  *SubscriptionRequest `json:"subscription"`
	Fields        []*NameValuePair     `json:"fields"`
	PaymentMethod null.String          `json:"paymentMethod"`
	Expiration    string               `json:"expiration"`
	ReturnURL     string               `json:"returnUrl"`
	IPAddres      string               `json:"ipAddress"`
	UserAgent     string               `json:"userAgent"`
	SkipResult    bool                 `json:"skipResult"`
	NoBuyerFill   bool                 `json:"noBuyerFill"`
}

// RedirectResponse structure
type RedirectResponse struct {
	Status     *Status `json:"status"`
	RequestID  int     `json:"requestID"`
	ProcessURL string  `json:"processUrl"`
}

// RedirectInformation structure
type RedirectInformation struct {
	Status       *Status               `json:"status"`
	Request      *RedirectRequest      `json:"request"`
	Payment      []*Transaction        `json:"payment"`
	Subscription *SubscriptionResponse `json:"subscription"`
}
