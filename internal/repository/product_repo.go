package repository

import (
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
)

type ProductDBRepositoryInterface interface {
	Create(product model.Product)(*model.Product, error)
	Get(id *int32)(*model.Product, error)
	GetAll()([]*model.Product, error)
	Delete(id *int32) error
	Update()
}

type ProductDBRepository struct {
	App *config.AppConfig

}
