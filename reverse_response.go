package placetopay

// ReverseResponse structure
type ReverseResponse struct {
	Status  *Status      `json:"status"`
	Payment *Transaction `json:"payment"`
}
