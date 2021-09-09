package repository

import "github.com/igor-koniukhov/fastcat/internal/config"

var Repo *Repository

type Repository struct{
	UserRepository
	SupplierRepository
	ProductRepository
	OrderRepository
	CartRepository
}

func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{
		UserRepository:     NewUserRepository(app),
		SupplierRepository: NewSupplierRepository(app),
		ProductRepository:  NewProductRepository(app),
		OrderRepository:    NewOrderRepository(app),
		CartRepository:     NewCartRepository(app),
	}
}

func NewRepo(r *Repository)  {
	Repo=r
}