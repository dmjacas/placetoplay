package placetopay

// SimpleToken structure
type SimpleToken struct {
	Token        string `json:"token"`
	Subtoken     string `json:"subtoken"`
	Installments int    `json:"installments"`
	Cvv          string `json:"cvv"`
}
