package repository

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"net/http"
)

type ProductRepositoryI interface {
	CreateProduct(u *model.Product, db *sql.DB) (*model.Product, error)
	GetProduct(nameParam, param *string, db *sql.DB) *model.Product
	GetAllUProduct(db *sql.DB) *[]model.Product
	DeleteProduct(id int, db *sql.DB) error
	UpdateProduct(id int, u *model.Product, db *sql.DB) *model.Product
	Param(r *http.Request) (string, string, int)
}

type ProductRepository struct {
	App *config.AppConfig

}
