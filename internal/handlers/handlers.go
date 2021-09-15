package handlers

import (
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/repository"
)

type Handlers struct {
	User
	Supplier
	Product
	Order
	Cart
}

func NewHandlers(app *config.AppConfig, repos *repository.Repository) *Handlers {
	return &Handlers{
		User:     NewUserHandler(app, repos.UserRepository),
		Supplier: NewSupplierHandler(app, repos.SupplierRepository),
		Product:  NewProductHandler(app, repos.ProductRepository),
		Order:    NewOrderHandler(app, repos.OrderRepository),
		Cart:     NewCartHandler(app, repos.CartRepository),
	}
}

