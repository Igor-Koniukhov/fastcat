package repository

import (
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"net/http"
)

type ProductRepositoryI interface {
	CreateProduct(u *model.Product) (*model.Product, error)
	GetProduct(nameParam, param *string) *model.Product
	GetAllUProducts() *[]model.Product
	DeleteProduct(id int) error
	UpdateProduct(id int, u *model.Product) *model.Product
	Param(r *http.Request) (string, string, int)
}

var RepoP *ProductRepository

type ProductRepository struct {
	App *config.AppConfig
}

func NewProductRepository(app *config.AppConfig) *ProductRepository {
	return &ProductRepository{App: app}
}
func NewRepoP(r *ProductRepository) {
	RepoP = r
}

func (p ProductRepository) CreateProduct(u *model.Product) (*model.Product, error) {
	return nil, nil
}

func (p ProductRepository) GetProduct(nameParam, param *string) *model.Product {
	return nil
}

func (p ProductRepository) GetAllUProducts() *[]model.Product {
	return nil
}

func (p ProductRepository) DeleteProduct(id int, ) error {
	return nil
}

func (p ProductRepository) UpdateProduct(id int, u *model.Product, ) *model.Product {
	return nil
}

func (p ProductRepository) Param(r *http.Request) (string, string, int) {
	return "", "", 0
}
