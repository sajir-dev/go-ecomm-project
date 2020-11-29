package domain

import (
	"fmt"

	"../db/config"
	uuid "github.com/google/uuid"
)

var m map[string]Product

// GetAllProducts returns all products from the db
func (p *Product) GetAllProducts() ([]Product, error) {
	// ToDo: Get all products query to DB
	var ps []Product
	// for _, v := range m {
	// 	ps = append(ps, v)
	// }

	rows, err := config.DB.Query(`SELECT * FROM products`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		var product Product
		rows.Scan(
			&product.ProductID,
			&product.ProductName,
			&product.Category,
			&product.Price,
			&product.Discount,
			&product.Description,
		)
		ps = append(ps, product)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return ps, nil
}

// CreateProduct creates a product in db
func (p *Product) CreateProduct(productName string, category string, price float32, discount float32, description string) (*Product, error) {
	// ToDo: insert into the db
	// insert into products (product_name, category, price, discount, description, rating) values ('test name','test cat', 34.5, 12, 'test description', 4.5);

	uuid := uuid.New().String()

	m[productName] = Product{
		ProductID:   uuid,
		Category:    category,
		Price:       price,
		Discount:    discount,
		Description: description,
	}

	product := m[productName]

	return &product, nil
}
