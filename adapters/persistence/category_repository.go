package persistence

import (
	"database/sql"
	"github.com/wevertonbruno/supermarket-api-go/core/domain"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (p CategoryRepository) Save(cat domain.Category) error {
	stmt, err := p.db.Prepare("INSERT INTO categories VALUES(?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(cat.ID, cat.Name, cat.Description)
	stmt.Close()
	return err
}

func (p CategoryRepository) Update(ID string, cat domain.Category) error {
	stmt, err := p.db.Prepare("UPDATE categories SET name = ?, description = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(cat.Name, cat.Description, ID)
	stmt.Close()
	return err
}
func (p CategoryRepository) FindAll() []domain.Category {
	rows, _ := p.db.Query("SELECT * FROM categories")

	var categories = make([]domain.Category, 0)

	i := 0
	for rows.Next() {
		var category domain.Category
		rows.Scan(&category.ID, &category.Name, &category.Description)
		categories = append(categories, category)
		i += 1
	}

	rows.Close()
	return categories
}
func (p CategoryRepository) FindById(ID string) (domain.Category, error) {
	row := p.db.QueryRow("SELECT * FROM categories WHERE id = ?", ID)

	var category domain.Category
	row.Scan(&category.ID, &category.Name, &category.Description)

	return category, nil
}
func (p CategoryRepository) Delete(ID string) error {
	stmt, err := p.db.Prepare("DELETE FROM categories WHERE id = ?")
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
