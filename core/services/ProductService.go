package services

import (
	"github.com/google/uuid"
	"github.com/wevertonbruno/supermarket-api-go/core/domain"
	"github.com/wevertonbruno/supermarket-api-go/core/ports"
	"github.com/wevertonbruno/supermarket-api-go/core/ports/DTO"
)

type ProductService struct {
	repository ports.ProductRepositoryPort
}

func NewProductService(repo ports.ProductRepositoryPort) *ProductService {
	return &ProductService{
		repository: repo,
	}
}

func (s ProductService) Save(dto DTO.InputProductDTO) error {
	prod := domain.Product{
		ID:       uuid.New().String(),
		Name:     dto.Name,
		Category: dto.Category,
		Price:    dto.Price,
	}

	return s.repository.Save(prod)
}

func (s ProductService) Update(ID string, dto DTO.InputProductDTO) error {
	prod, err := s.repository.FindById(ID)
	if err != nil {
		return err
	}

	prod.Name = dto.Name
	prod.Category = dto.Category
	prod.Price = dto.Price

	return s.repository.Update(ID, prod)
}

func (s ProductService) FindAll() []DTO.OutputProductDTO {
	data := s.repository.FindAll()
	dto := make([]DTO.OutputProductDTO, len(data))

	for k, v := range data {
		dto[k] = DTO.OutputProductDTO{
			ID:       v.ID,
			Name:     v.Name,
			Category: v.Category,
			Price:    v.Price,
		}
	}

	return dto
}

func (s ProductService) FindById(ID string) (DTO.OutputProductDTO, error) {
	v, err := s.repository.FindById(ID)
	if err != nil {
		return DTO.OutputProductDTO{}, err
	}

	dto := DTO.OutputProductDTO{
		ID:       v.ID,
		Name:     v.Name,
		Category: v.Category,
		Price:    v.Price,
	}

	return dto, nil
}

func (s ProductService) Delete(ID string) error {
	_, err := s.FindById(ID)
	if err != nil {
		return err
	}

	return s.repository.Delete(ID)
}
