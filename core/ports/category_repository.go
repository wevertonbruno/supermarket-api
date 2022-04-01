package ports

import "github.com/wevertonbruno/supermarket-api-go/core/domain"

type CategoryRepositoryPort interface {
	Save(p domain.Category) error
	Update(ID string, product domain.Category) error
	FindAll() []domain.Category
	FindById(ID string) (domain.Category, error)
	Delete(ID string) error
}
