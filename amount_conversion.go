package placetopay

//AmountConversion structure
type AmountConversion struct {
	From   *AmountBase `json:"from"`
	To     *AmountBase `json:"to"`
	Factor float64     `json:"factor"`
}
