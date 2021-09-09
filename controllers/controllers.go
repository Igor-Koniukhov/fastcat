package controllers

import "github.com/igor-koniukhov/fastcat/internal/repository"

type Controllers struct {
	User
	Supplier
	Product
	Order
	Cart
}

var Controller *Controllers

func NewControllers(repos *repository.Repository) *Controllers {
	return &Controllers{
		User:     NewUserController(repos.UserRepository),
		Supplier: NewSupplierController(repos.SupplierRepository),
		Product:  NewProductController(repos.ProductRepository),
		Order:    NewOrderController(repos.OrderRepository),
		Cart:     NewCartController(repos.CartRepository),
	}
}

func NewController(c *Controllers){
	Controller = c
}