package services

import (
	"github.com/google/uuid"
	"github.com/wevertonbruno/supermarket-api-go/core/domain"
	"github.com/wevertonbruno/supermarket-api-go/core/ports"
	"github.com/wevertonbruno/supermarket-api-go/core/ports/DTO"
)

type CategoryService struct {
	repository ports.CategoryRepositoryPort
}

func NewCategoryService(repo ports.CategoryRepositoryPort) *CategoryService {
	return &CategoryService{
		repository: repo,
	}
}

func (s CategoryService) Save(dto DTO.InputCategoryDTO) error {
	cat := domain.Category{
		ID:          uuid.New().String(),
		Name:        dto.Name,
		Description: dto.Description,
	}

	return s.repository.Save(cat)
}

func (s CategoryService) Update(ID string, dto DTO.InputCategoryDTO) error {
	cat, err := s.repository.FindById(ID)
	if err != nil {
		return err
	}

	cat.Name = dto.Name
	cat.Description = dto.Description

	return s.repository.Update(ID, cat)
}

func (s CategoryService) FindAll() []DTO.OutputCategoryDTO {
	data := s.repository.FindAll()
	dto := make([]DTO.OutputCategoryDTO, len(data))

	for k, v := range data {
		dto[k] = DTO.OutputCategoryDTO{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
		}
	}

	return dto
}

func (s CategoryService) FindById(ID string) (DTO.OutputCategoryDTO, error) {
	v, err := s.repository.FindById(ID)
	if err != nil {
		return DTO.OutputCategoryDTO{}, err
	}

	dto := DTO.OutputCategoryDTO{
		ID:          v.ID,
		Name:        v.Name,
		Description: v.Description,
	}

	return dto, nil
}

func (s CategoryService) Delete(ID string) error {
	_, err := s.FindById(ID)
	if err != nil {
		return err
	}

	return s.repository.Delete(ID)
}
