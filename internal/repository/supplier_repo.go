package repository

import (
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
)

type SupplierRepositoryInterface interface {
	Create(supplier model.Supplier) (*model.Supplier, error)
	Get(id *int32) (*model.Supplier, error)
	GetAll() ([]*model.Supplier, error)
	Delete(id *int32) error
}

type SupplierDBRepository struct {
	App *config.AppConfig
}

func NewSupplierDBRepository(a *config.AppConfig) *SupplierDBRepository {
	return &SupplierDBRepository{
		App: a,
	}
}

func (s SupplierDBRepository) Create(supplier model.Supplier) (*model.Supplier, error) {
return nil, nil
}

func (s SupplierDBRepository) Get(id *int32) (*model.Supplier, error) {
	return nil, nil


}

func (s SupplierDBRepository) GetAll() ([]*model.Supplier, error) {
	return nil, nil

}

func (s SupplierDBRepository) Delete(id *int32) error {
return nil
}
