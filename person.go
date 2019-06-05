package placetopay

//Person structure
type Person struct {
	DocumenType string  `json:"documentType"`
	Document    string  `json:"document"`
	Name        string  `json:"name"`
	Surname     string  `json:"surname"`
	Company     string  `json:"company"`
	Email       string  `json:"email"`
	Addres      *Addres `json:"addres"`
	Mobile      string  `json:"mobile"`
}
