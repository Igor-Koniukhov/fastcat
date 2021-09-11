package handlers

import "github.com/igor-koniukhov/fastcat/internal/repository"

type Handlers struct {
	User
	Supplier
	Product
	Order
	Cart
}

func NewHandlers(repos *repository.Repository) *Handlers {
	return &Handlers{
		User:     NewUserHandler(repos.UserRepository),
		Supplier: NewSupplierHandler(repos.SupplierRepository),
		Product:  NewProductHandler(repos.ProductRepository),
		Order:    NewOrderHandler(repos.OrderRepository),
		Cart:     NewCartHandler(repos.CartRepository),
	}
}

