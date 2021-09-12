package main

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/handlers"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"net/http"
)

func routes(app *config.AppConfig, db *sql.DB) http.Handler {

	repo := repository.NewRepository(app, db)
	www := handlers.NewHandlers(repo)
	repository.NewRepo(repo)

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./public")))
	mux.HandleFunc("/login",	AuthMiddleWare(www.User.Login))
	mux.HandleFunc("/refresh", www.User.Refresh)
	mux.HandleFunc("/logout", www.User.LogOut)

	mux.HandleFunc("/user/create", www.User.Create)
	mux.HandleFunc("/user/", www.User.Get)
	mux.HandleFunc("/users", www.User.GetAll)
	mux.HandleFunc("/user/update/", www.User.Update)
	mux.HandleFunc("/user/delete/", www.User.Delete)

	mux.HandleFunc("/order/create", www.Order.Create)
	mux.HandleFunc("/order/", www.Order.Get)
	mux.HandleFunc("/orders", www.Order.GetAll)
	mux.HandleFunc("/order/update/", www.Order.Update)
	mux.HandleFunc("/order/delete/", www.Order.Delete)

	mux.HandleFunc("/supplier/create", www.Supplier.Create)
	mux.HandleFunc("/supplier/", www.Supplier.Get)
	mux.HandleFunc("/suppliers", www.Supplier.GetAll)
	mux.HandleFunc("/supplier/update/", www.Supplier.Update)
	mux.HandleFunc("/supplier/delete/", www.Supplier.Delete)

	mux.HandleFunc("/product/create", www.Product.Update)
	mux.HandleFunc("/product/", www.Product.Get)
	mux.HandleFunc("/products", www.Product.GetAll)
	mux.HandleFunc("/product/update/", www.Product.Update)
	mux.HandleFunc("/product/delete/", www.Product.Delete)

	mux.HandleFunc("/cart/create", www.Cart.Create)
	mux.HandleFunc("/cart/", www.Cart.Get)
	mux.HandleFunc("/cart", www.Cart.GetAll)
	mux.HandleFunc("/cart/update/", www.Cart.Update)
	mux.HandleFunc("/cart/delete/", www.Cart.Delete)

	return mux
}
