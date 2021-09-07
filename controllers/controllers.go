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
		User:     NewUserController(repos.UserRepositoryInterface),
		Supplier: NewSupplierController(repos.SupplierRepositoryInterface),
		Product:  NewProductController(repos.ProductRepositoryInterface),
		Order:    NewOrderController(repos.OrderRepositoryInterface),
		Cart:     NewCartController(repos.CartRepositoryInterface),
	}
}

func NewController(c *Controllers){
	Controller = c
}