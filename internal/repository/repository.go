package repository

import (
	"github.com/igor-koniukhov/fastcat/driver"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
)

var Repo *Repository

type Repository struct{
	dbrepo.UserRepository
	dbrepo.SupplierRepository
	dbrepo.ProductRepository
	dbrepo.OrderRepository
	dbrepo.CartRepository
}

func NewRepository(app *config.AppConfig, DB *driver.DB) *Repository {
	return &Repository{
		UserRepository:     dbrepo.NewUserRepository(app, DB.MySQL),
		SupplierRepository: dbrepo.NewSupplierRepository(app, DB.MySQL),
		ProductRepository:  dbrepo.NewProductRepository(app, DB.MySQL),
		OrderRepository:    dbrepo.NewOrderRepository(app, DB.MySQL),
		CartRepository:     dbrepo.NewCartRepository(app, DB.MySQL),
	}
}

func NewRepo(r *Repository)  {
	Repo=r
}