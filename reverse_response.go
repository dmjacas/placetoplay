package placetopay

// ReverseResponse structure
type ReverseResponse struct {
	Status  *Status      `json:"status"`
	Payment *Transaction `json:"payment"`
}

// ReverseBodyRequest structure
type ReverseBodyRequest struct {
	Auth              *Auth  `json:"status"`
	InternalReference string `json:"internalReference"`
}
