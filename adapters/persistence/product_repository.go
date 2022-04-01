package persistence

import (
	"database/sql"
	"github.com/wevertonbruno/supermarket-api-go/core/domain"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p ProductRepository) Save(prod domain.Product) error {
	stmt, err := p.db.Prepare("INSERT INTO products VALUES(?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(prod.ID, prod.Name, prod.Category.ID, prod.Price)
	stmt.Close()
	return err
}

func (p ProductRepository) Update(ID string, prod domain.Product) error {
	stmt, err := p.db.Prepare("UPDATE products SET name = ?, category_id = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(prod.Name, prod.Category.ID, prod.Price, ID)
	stmt.Close()
	return err
}
func (p ProductRepository) FindAll() []domain.Product {
	rows, _ := p.db.Query("SELECT * FROM products")
	defer rows.Close()

	var products = make([]domain.Product, 0)
	var categories = map[string]domain.Category{}

	rowsCategory, _ := p.db.Query("SELECT * FROM categories")
	defer rowsCategory.Close()

	for rowsCategory.Next() {
		var category domain.Category
		rowsCategory.Scan(&category.ID, &category.Name, &category.Description)
		categories[category.ID] = category
	}

	for rows.Next() {
		var product domain.Product
		rows.Scan(&product.ID, &product.Name, &product.Category.ID, &product.Price)
		product.Category = categories[product.Category.ID]
		products = append(products, product)
	}

	return products
}
func (p ProductRepository) FindById(ID string) (domain.Product, error) {
	row := p.db.QueryRow("SELECT * FROM products WHERE id = ?", ID)

	var product domain.Product
	row.Scan(&product.ID, &product.Name, &product.Category.ID, &product.Price)

	rowCategory := p.db.QueryRow("SELECT * FROM categories WHERE id = ?", product.Category.ID)
	var category domain.Category
	rowCategory.Scan(&category.ID, &category.Name, &category.Description)
	product.Category = category

	return product, nil
}
func (p ProductRepository) Delete(ID string) error {
	stmt, err := p.db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(ID)
	stmt.Close()

	if err != nil {
		return err
	}

	return nil
}
