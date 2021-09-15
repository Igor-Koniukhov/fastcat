package main

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/handlers"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"github.com/igor-koniukhov/fastcat/middleware"
	"net/http"
)

func routes(app *config.AppConfig, db *sql.DB) http.Handler {
	repo := repository.NewRepository(app, db)
	www := handlers.NewHandlers(app, repo)
	repository.NewRepo(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/login", www.User.PostLogin)
	mux.HandleFunc("/show-login", www.User.ShowLogin)
	mux.HandleFunc("/refresh", www.User.Refresh)
	mux.HandleFunc("/logout", www.User.Logout)

	mux.HandleFunc("/user/create", www.User.Create)
	mux.HandleFunc("/user/", middleware.Auth(www.User.Get))
	mux.HandleFunc("/users", www.User.GetAll)
	mux.HandleFunc("/user/update/", middleware.Auth(www.User.Update))
	mux.HandleFunc("/user/delete/", middleware.Auth(www.User.Delete))

	mux.HandleFunc("/order/create", middleware.Auth(www.Order.Create))
	mux.HandleFunc("/order/", middleware.Auth(www.Order.Get))
	mux.HandleFunc("/orders", middleware.Auth(www.Order.GetAll))
	mux.HandleFunc("/order/update/", middleware.Auth(www.Order.Update))
	mux.HandleFunc("/order/delete/", middleware.Auth(www.Order.Delete))

	mux.HandleFunc("/supplier/create", middleware.Auth(www.Supplier.Create))
	mux.HandleFunc("/supplier/", www.Supplier.Get)
	mux.HandleFunc("/suppliers", www.Supplier.GetAll)
	mux.HandleFunc("/supplier/update/", middleware.Auth(www.Supplier.Update))
	mux.HandleFunc("/supplier/delete/", middleware.Auth(www.Supplier.Delete))

	mux.HandleFunc("/product/create", middleware.Auth(www.Product.Update))
	mux.HandleFunc("/product/", www.Product.Get)
	mux.HandleFunc("/products", www.Product.GetAll)
	mux.HandleFunc("/product/update/", middleware.Auth(www.Product.Update))
	mux.HandleFunc("/product/delete/", middleware.Auth(www.Product.Delete))

	mux.HandleFunc("/cart/create", www.Cart.Create)
	mux.HandleFunc("/cart/", www.Cart.Get)
	mux.HandleFunc("/carts", www.Cart.GetAll)
	mux.HandleFunc("/cart/update/", middleware.Auth(www.Cart.Update))
	mux.HandleFunc("/cart/delete/", middleware.Auth(www.Cart.Delete))

	fileServe := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServe))

	return mux
}
