package ports

import (
	"github.com/wevertonbruno/supermarket-api-go/core/ports/DTO"
)

type ProductServicePort interface {
	Save(p DTO.InputProductDTO) error
	Update(ID string, dto DTO.InputProductDTO) error
	FindAll() []DTO.OutputProductDTO
	FindById(ID string) (DTO.OutputProductDTO, error)
	Delete(ID string) error
}
