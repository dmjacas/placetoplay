package placetopay

import "time"

// Status structure
type Status struct {
	Status  string    `json:"status"`
	Reason  string    `json:"reason"`
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}
