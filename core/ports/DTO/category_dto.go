package DTO

type InputCategoryDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type OutputCategoryDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
