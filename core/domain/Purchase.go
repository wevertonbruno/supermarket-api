package domain

type Purchase struct {
	ID    string        `json:"id"`
	Items []ItemProduct `json:"items"`
	total float32       `json:"total"`
}
