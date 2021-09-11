package repository

import (
	"database/sql"
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

func NewRepository(app *config.AppConfig, db *sql.DB) *Repository {
	return &Repository{
		UserRepository:     dbrepo.NewUserRepository(app, db),
		SupplierRepository: dbrepo.NewSupplierRepository(app, db),
		ProductRepository:  dbrepo.NewProductRepository(app, db),
		OrderRepository:    dbrepo.NewOrderRepository(app, db),
		CartRepository:     dbrepo.NewCartRepository(app, db),
	}
}

func NewRepo(r *Repository)  {
	Repo=r
}