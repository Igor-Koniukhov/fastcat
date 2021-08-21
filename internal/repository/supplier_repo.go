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

func (s SupplierRepository) CreateSupplier(u *model.Supplier, db *sql.DB) (*model.Supplier, error) {
	return nil, nil

}

func (s SupplierRepository) GetSupplier(nameParam, param *string, db *sql.DB) *model.Supplier {
	return nil
}

func (s SupplierRepository) GetAllUSuppliers(db *sql.DB) *[]model.Supplier {
	return nil
}

func (s SupplierRepository) DeleteSupplier(id int, db *sql.DB) error {
	return nil
}

func (s SupplierRepository) UpdateSupplier(id int, u *model.Supplier, db *sql.DB) *model.Supplier {
	return nil
}

func (s SupplierRepository) Param(r *http.Request) (string, string, int) {
	return "", "", 0
}

