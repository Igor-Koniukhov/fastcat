package repository

import (
	"github.com/igor-koniukhov/fastcat/internal/model"
	"net/http"
)

type SupplierRepositoryI interface {
	CreateSupplier(u *model.Supplier) (*model.Supplier, error)
	GetSupplier(nameParam, param *string, ) *model.Supplier
	GetAllUSuppliers() *[]model.Supplier
	DeleteSupplier(id int, ) error
	UpdateSupplier(id int, u *model.Supplier) *model.Supplier
	Param(r *http.Request) (string, string, int)
}

type SupplierRepository struct {}

func (s SupplierRepository) CreateSupplier(u *model.Supplier, ) (*model.Supplier, error) {
	return nil, nil

}

func (s SupplierRepository) GetSupplier(nameParam, param *string, ) *model.Supplier {
	return nil
}

func (s SupplierRepository) GetAllUSuppliers() *[]model.Supplier {
	return nil
}

func (s SupplierRepository) DeleteSupplier(id int, ) error {
	return nil
}

func (s SupplierRepository) UpdateSupplier(id int, u *model.Supplier, ) *model.Supplier {
	return nil
}

func (s SupplierRepository) Param(r *http.Request) (string, string, int) {
	return "", "", 0
}
