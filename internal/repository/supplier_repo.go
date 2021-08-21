package repository

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"net/http"
)

type SupplierRepositoryI interface {
	CreateSupplier(u *model.Supplier, db *sql.DB) (*model.Supplier, error)
	GetSupplier(nameParam, param *string, db *sql.DB) *model.Supplier
	GetAllUSuppliers(db *sql.DB) *[]model.Supplier
	DeleteSupplier(id int, db *sql.DB) error
	UpdateSupplier(id int, u *model.Supplier, db *sql.DB) *model.Supplier
	Param(r *http.Request) (string, string, int)
}

type SupplierRepository struct {
	App *config.AppConfig
}
