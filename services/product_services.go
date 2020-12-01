package services

import (
	"../domain"
)

type productServicesInterface interface {
	CreateProduct(string, string, float32, float32, string, float32) (*domain.Product, error)
	GetAllProducts() ([]domain.Product, error)
	GetProduct(string) (*domain.Product, error)
	UpdateProduct(string, string, string, float32, float32, string, float32) (*domain.Product, error)
}

type productServices struct{}

var (
	ProductServices productServicesInterface = &productServices{}
)

func (ps *productServices) CreateProduct(productName string, category string, price float32, discount float32, description string, rating float32) (*domain.Product, error) {
	p, err := domain.CreateProduct(productName, category, price, discount, description, rating)
	return p, err
}

func (ps *productServices) GetAllProducts() ([]domain.Product, error) {
	p, err := domain.GetAllProducts()
	return p, err
}

func (ps *productServices) GetProduct(id string) (*domain.Product, error) {
	p, err := domain.GetProduct(id)
	return p, err
}

func (ps *productServices) UpdateProduct(id string, productName string, category string, price float32, discount float32, description string, rating float32) (*domain.Product, error) {
	p, err := domain.UpdateProduct(id, productName, category, price, discount, description, rating)
	return p, err
}
