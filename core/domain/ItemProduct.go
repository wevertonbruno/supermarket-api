package domain

type ItemProduct struct {
	Product  Product `json:"product"`
	Quantity uint    `json:"quantity"`
	SubTotal float32 `json:"subTotal"`
}
