package persistence

import (
	"database/sql"
)

func CreateDb(db *sql.DB) error {
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS categories(
    	id varchar not null PRIMARY KEY ,
    	name varchar not null,
    	description text not null
	)`)

	if err != nil {
		return err
	}

	stmt.Exec()
	stmt.Close()

	//initialize table of products
	stmt, err = db.Prepare(`CREATE TABLE IF NOT EXISTS products (
    	id varchar not null PRIMARY KEY ,
    	name varchar not null,
    	category_id varchar not null,
    	price decimal(10,2) not null,
    	foreign key (category_id) references category(id)
	)`)

	if err != nil {
		return err
	}

	stmt.Exec()
	stmt.Close()

	return nil
}
