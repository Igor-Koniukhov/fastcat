package repository

import (
	"github.com/igor-koniukhov/fastcat/internal/model"
	"net/http"
)

type ProductRepositoryI interface {
	CreateProduct(u *model.Product ) (*model.Product, error)
	GetProduct(nameParam, param *string) *model.Product
	GetAllUProducts() *[]model.Product
	DeleteProduct(id int ) error
	UpdateProduct(id int, u *model.Product ) *model.Product
	Param(r *http.Request) (string, string, int)
}

type ProductRepository struct {}

func (p ProductRepository) CreateProduct(u *model.Product) (*model.Product, error) {
	return nil, nil
}


func (p ProductRepository) GetProduct(nameParam, param *string) *model.Product {
	return nil
}

func (p ProductRepository) GetAllUProducts() *[]model.Product {
	return nil
}

func (p ProductRepository) DeleteProduct(id int,) error {
	return nil
}

func (p ProductRepository) UpdateProduct(id int, u *model.Product,) *model.Product {
	return nil
}

func (p ProductRepository) Param(r *http.Request) (string, string, int) {
	return "", "", 0
}

