package persistence

import "github.com/wevertonbruno/supermarket-api-go/core/domain"

type ProductRepository struct {
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (p ProductRepository) Save(prod domain.Product) error {
	return nil
}

func (p ProductRepository) Update(ID string, product domain.Product) error {
	return nil
}
func (p ProductRepository) FindAll() []domain.Product {
	return []domain.Product{}
}
func (p ProductRepository) FindById(ID string) (domain.Product, error) {
	return domain.Product{}, nil
}
func (p ProductRepository) Delete(ID string) error {
	return nil
}
