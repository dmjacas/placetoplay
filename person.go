package placetopay

import "gopkg.in/guregu/null.v3"

//Person structure
type Person struct {
	DocumenType string      `json:"documentType"`
	Document    string      `json:"document"`
	Name        string      `json:"name"`
	Surname     string      `json:"surname"`
	Company     null.String `json:"company"`
	Email       string      `json:"email"`
	Addres      *Addres     `json:"addres"`
	Mobile      string      `json:"mobile"`
}
