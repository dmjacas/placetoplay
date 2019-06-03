package placetopay

//Item structure
type Item struct {
	Sku      string  `json:"street"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Qty      string  `json:"qty"`
	Price    float64 `json:"price"`
	Tax      float64 `json:"tax"`
}
