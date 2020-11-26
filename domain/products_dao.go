package domain

import (
	uuid "github.com/google/uuid"
)

// Product struct
type Product struct {
	ProductID   string  `json:"id"`
	ProductName string  `json:"product_name"`
	Category    string  `json:"category"`
	Price       float32 `json:"price"`
	Discount    float32 `json:"discout"`
	Description string  `json:"description"`
	// Rating      float32 `json:"rating"`
	// Stock       int     `json:"stock"`
	// Images      string  `json:"images"`
	// Size        string  `json:"size"`
}

var m map[string]Product

// GetProducts returns all products from the db
func (p *Product) GetProducts() ([]Product, error) {
	// ToDo: Get all products query to DB
	var ps []Product
	for _, v := range m {
		ps = append(ps, v)
	}
	return ps, nil
}

// CreateProduct creates a product in db
func (p *Product) CreateProduct(productName string, category string, price float32, discount float32, description string) (*Product, error) {
	// ToDo: insert into the db

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
