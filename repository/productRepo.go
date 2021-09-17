package repository

import "go-mongo/models"

type ProductRepository interface {
	FindAllProduct() ([]models.Product, error)
	CreateProduct(product models.Product) error
	GetProductByName(name string) (models.Product, error)
}
