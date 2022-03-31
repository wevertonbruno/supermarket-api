package domain

type Product struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Category Category `json:"category"`
	Price    float32  `json:"price"`
}
