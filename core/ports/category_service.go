package ports

import (
	"github.com/wevertonbruno/supermarket-api-go/core/ports/DTO"
)

type CategoryServicePort interface {
	Save(p DTO.InputCategoryDTO) error
	Update(ID string, dto DTO.InputCategoryDTO) error
	FindAll() []DTO.OutputCategoryDTO
	FindById(ID string) (DTO.OutputCategoryDTO, error)
	Delete(ID string) error
}
