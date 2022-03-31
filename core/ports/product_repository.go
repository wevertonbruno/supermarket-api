package ports

import "github.com/wevertonbruno/supermarket-api-go/core/domain"

type ProductRepositoryPort interface {
	Save(p domain.Product) error
	Update(ID string, product domain.Product) error
	FindAll() []domain.Product
	FindById(ID string) (domain.Product, error)
	Delete(ID string) error
}
