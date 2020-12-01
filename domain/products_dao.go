package domain

import (
	"fmt"

	"../db/config"
)

// var m map[string]Product

// GetAllProducts returns all products from the db
func GetAllProducts() ([]Product, error) {
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
			&product.Rating,
		)
		ps = append(ps, product)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return ps, nil
}

// GetProduct returns a single product based on the input given
func GetProduct(id string) (*Product, error) {
	var p *Product
	row := config.DB.QueryRow(`select * from products where id = "` + id + `"`)
	err := row.Scan(&p.ProductID, &p.ProductName, &p.Category, &p.Price, &p.Discount, &p.Description, &p.Rating)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return p, nil
}

// CreateProduct creates a product in db
func CreateProduct(productName string, category string, price float32, discount float32, description string, rating float32) (*Product, error) {
	// ToDo: insert into the db
	// insert into products (product_name, category, price, discount, description, rating) values ('test name','test cat', 34.5, 12, 'test description', 4.5);

	stmt, err := config.DB.Exec(`insert into products (product_name, category, price, discount, description, rating) values ($1, $2, $3, $4, $5, $6)`,
		productName,
		category,
		price,
		discount,
		description,
		rating)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	id, _ := stmt.LastInsertId()
	row := config.DB.QueryRow(`select * from products where id = "` + fmt.Sprint(id) + `"`)
	var product *Product
	err = row.Scan(&product.ProductID, &product.ProductName, &product.Category, &product.Price, &product.Description, &product.Rating)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return product, nil
	// uuid := uuid.New().String()

	// m[productName] = Product{
	// 	ProductID:   uuid,
	// 	Category:    category,
	// 	Price:       price,
	// 	Discount:    discount,
	// 	Description: description,
	// }

	// product := m[productName]

	// return &product, nil
}

// UpdateProduct updates the given arguments of a particular product
func UpdateProduct(id string, productName string, category string, price float32, discount float32, description string, rating float32) (*Product, error) {

	row := config.DB.QueryRow(`select * from products where id = "` + fmt.Sprint(id) + `"`)
	var product *Product
	err := row.Scan(&product.ProductID, &product.ProductName, &product.Category, &product.Price, &product.Description, &product.Rating)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = config.DB.Exec(`insert into products (product_name, category, price, discount, description, rating) values ($1, $2, $3, $4, $5, $6) where id=`+id,
		productName,
		category,
		price,
		discount,
		description,
		rating)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	row = config.DB.QueryRow(`select * from products where id = "` + fmt.Sprint(id) + `"`)
	err = row.Scan(&product.ProductID, &product.ProductName, &product.Category, &product.Price, &product.Description, &product.Rating)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return product, nil
}
