package placetopay

import "time"

//Recurring structure
type Recurring struct {
	Periodicity     string    `json:"periodicity"`
	Interval        int       `json:"interval"`
	NextPayment     time.Time `json:"nextPayment"`
	MaxPeriods      int       `json:"maxPeriods"`
	DueDate         time.Time `json:"dueDate"`
	NotificationURL string    `json:"notificationUrl"`
}
