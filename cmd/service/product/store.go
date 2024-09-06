package product

import (
	"database/sql"

	"github.com/femilshajin/ecomgo/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) AddProduct(product types.Product) (int64, error) {
	r, err := s.db.Exec("INSERT INTO products (name, description, image, price, quantity) VALUES ( ?,?,?,?,?)",
		product.Name, product.Description, product.Image, product.Price, product.Quantity)
	if err != nil {
		return -1, err
	}
	id, _ := r.LastInsertId()
	return id, nil
}

func (s *Store) GetProducts() ([]*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]*types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}
