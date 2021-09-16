package main

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/handlers"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	m "github.com/igor-koniukhov/fastcat/middleware"
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
	mux.Handle("/user/", m.Auth(http.HandlerFunc(www.User.Get)))
	mux.HandleFunc("/users", www.User.GetAll)
	mux.Handle("/user/update/", m.Auth(http.HandlerFunc(www.User.Update)))
	mux.Handle("/user/delete/", m.Auth(http.HandlerFunc(www.User.Delete)))

	mux.Handle("/order/create", m.Auth(http.HandlerFunc(www.Order.Create)))
	mux.Handle("/order/", m.Auth(http.HandlerFunc(www.Order.Get)))
	mux.Handle("/orders", m.Auth(http.HandlerFunc(www.Order.GetAll)))
	mux.Handle("/order/update/", m.Auth(http.HandlerFunc(www.Order.Update)))

	mux.Handle("/supplier/create", m.Auth(http.HandlerFunc(www.Supplier.Create)))
	mux.HandleFunc("/supplier/", www.Supplier.Get)
	mux.HandleFunc("/suppliers", www.Supplier.GetAll)
	mux.Handle("/supplier/update/", m.Auth(http.HandlerFunc(www.Supplier.Update)))
	mux.Handle("/supplier/delete/", m.Auth(http.HandlerFunc(www.Supplier.Delete)))

	mux.Handle("/product/create", m.Auth(http.HandlerFunc(www.Product.Update)))
	mux.HandleFunc("/product/", www.Product.Get)
	mux.HandleFunc("/products", www.Product.GetAll)
	mux.Handle("/product/update/", m.Auth(http.HandlerFunc(www.Product.Update)))
	mux.Handle("/product/delete/", m.Auth(http.HandlerFunc(www.Product.Delete)))

	mux.HandleFunc("/cart/create", www.Cart.Create)
	mux.HandleFunc("/cart/", www.Cart.Get)
	mux.HandleFunc("/carts", www.Cart.GetAll)
	mux.Handle("/cart/update/", m.Auth(http.HandlerFunc(www.Cart.Update)))
	mux.Handle("/cart/delete/", m.Auth(http.HandlerFunc(www.Cart.Delete)))

	fileServe := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServe))

	return mux
}

