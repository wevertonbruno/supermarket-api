package DTO

import "github.com/wevertonbruno/supermarket-api-go/core/domain"

type InputProductDTO struct {
	Name     string          `json:"name"`
	Category domain.Category `json:"category"`
	Price    float32         `json:"price"`
}

type OutputProductDTO struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Category domain.Category `json:"category"`
	Price    float32         `json:"price"`
}
