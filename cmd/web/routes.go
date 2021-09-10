package main

import (
	"github.com/igor-koniukhov/fastcat/controllers"
	"github.com/igor-koniukhov/fastcat/driver"
	"github.com/igor-koniukhov/fastcat/internal/auth/handlers"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"net/http"
)

func routes(app *config.AppConfig, db *driver.DB) http.Handler {

	repo := repository.NewRepository(app, db)
	c := controllers.NewControllers(repo)
	repository.NewRepo(repo)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./public")))
	mux.HandleFunc("/login",AuthMiddleWare(handlers.Login))
	mux.HandleFunc("/refresh",handlers.Refresh)

	mux.HandleFunc("/user/create", c.User.Create())
	mux.HandleFunc("/user/", c.User.Get())
	mux.HandleFunc("/users", c.User.GetAll())
	mux.HandleFunc("/user/update/", c.User.Update())
	mux.HandleFunc("/user/delete/", c.User.Delete())

	mux.HandleFunc("/order/create", c.Order.Create())
	mux.HandleFunc("/order/", c.Order.Get())
	mux.HandleFunc("/orders", c.Order.GetAll())
	mux.HandleFunc("/order/update/", c.Order.Update())
	mux.HandleFunc("/order/delete/", c.Order.Delete())

	mux.HandleFunc("/supplier/create", c.Supplier.Create())
	mux.HandleFunc("/supplier/", c.Supplier.Get())
	mux.HandleFunc("/suppliers", c.Supplier.GetAll())
	mux.HandleFunc("/supplier/update/", c.Supplier.Update())
	mux.HandleFunc("/supplier/delete/", c.Supplier.Delete())

	mux.HandleFunc("/product/create", c.Product.Update())
	mux.HandleFunc("/product/", c.Product.Get())
	mux.HandleFunc("/products", c.Product.GetAll())
	mux.HandleFunc("/product/update/", c.Product.Update())
	mux.HandleFunc("/product/delete/", c.Product.Delete())

	mux.HandleFunc("/cart/create", c.Cart.Create())
	mux.HandleFunc("/cart/", c.Cart.Get())
	mux.HandleFunc("/cart", c.Cart.GetAll())
	mux.HandleFunc("/cart/update/", c.Cart.Update())
	mux.HandleFunc("/cart/delete/", c.Cart.Delete())

	return mux
}
