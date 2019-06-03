package placetopay

//SubscriptionRequest structure
type SubscriptionRequest struct {
	Reference   string         `json:"reference"`
	Description string         `json:"description"`
	Fields      *NameValuePair `json:"fields"`
}

//SubscriptionResponse structure
type SubscriptionResponse struct {
	Status    *Status        `json:"reference"`
	Type      string         `json:"type"`
	Intrument *NameValuePair `json:"fields"`
}
