package repository

import "github.com/igor-koniukhov/fastcat/internal/config"

type Repository struct{
	UserRepositoryInterface
	SupplierRepositoryInterface
	ProductRepositoryInterface
	OrderRepositoryInterface
	CartRepositoryInterface
}

func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{
		UserRepositoryInterface:     NewUserRepository(app),
		SupplierRepositoryInterface: NewSupplierRepository(app),
		ProductRepositoryInterface:  NewProductRepository(app),
		OrderRepositoryInterface:    NewOrderRepository(app),
		CartRepositoryInterface:     NewCartRepository(app),
	}
}