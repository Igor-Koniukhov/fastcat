package main

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/handlers"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	m "github.com/igor-koniukhov/fastcat/middleware"
	"github.com/igor-koniukhov/fastcat/services/router"
	"net/http"
)

func routes(app *config.AppConfig, db *sql.DB) http.Handler {
	repo := repository.NewRepository(app, db)
	www := handlers.NewHandlers(app, repo)
	repository.NewRepo(repo)

	mux := router.NewRoutServeMux()
	mux.GET("/", mux.CORS(www.Supplier.Home))
	mux.GET("/registration", mux.CORS(www.User.ShowRegistration))
	mux.POST("/sign-up", mux.CORS(www.User.SingUp))
	mux.GET("/show-login", www.User.ShowLogin)
	mux.POST("/login", mux.CORS(www.User.PostLogin))
	mux.GET("/refresh", www.User.RefreshToken)
	mux.GET("/logout", www.User.Logout)

	mux.GET("/about", www.User.AboutUs)
	mux.GET("/status", m.Auth(www.User.StatusPage))
	mux.GET("/contacts", www.User.Contacts)
	mux.GET("/user/:id", mux.CORS(m.Auth(www.User.Get)))
	mux.GET("/users", www.User.GetAll)
	mux.PUT("/user/update/:id", mux.CORS(m.Auth(www.User.Update)))
	mux.DEL("/user/delete/:id", mux.CORS(m.Auth(www.User.Delete)))

	mux.GET("/order/blank", m.Auth(www.Order.ShowBlankOrder))
	mux.POST("/order/create", mux.CORS(m.Auth(www.Order.Create)))
	mux.GET("/order/:id", mux.CORS(m.Auth(www.Order.Get)))
	mux.GET("/orders", mux.CORS(m.Auth(www.Order.GetAll)))
	mux.PUT("/order/update/:id", mux.CORS(m.Auth(www.Order.Update)))

	mux.GET("/supplier/:id", m.Auth(www.Supplier.Get))
	mux.GET("/suppliers", m.Auth(www.Supplier.GetAll))
	mux.GET("/suppliers/selected", m.Auth(www.Supplier.GetAllBySchedule))
	mux.PUT("/supplier/update/:id", mux.CORS(m.Auth(www.Supplier.Update)))
	mux.DEL("/supplier/delete/:id", mux.CORS(m.Auth(www.Supplier.Delete)))

	mux.GET("/product/:id", m.Auth(www.Product.Get))
	mux.GET("/suppliers-products", m.Auth(www.Product.GetAllBySupplierID))
	mux.GET("/products", m.Auth(www.Product.GetAll))
	mux.GET("/products-json", mux.JSON(www.Product.GetJson))
	mux.PUT("/product/update/:id", mux.CORS(m.Auth(www.Product.Update)))
	mux.DEL("/product/delete/:id", mux.CORS(m.Auth(www.Product.Delete)))

	mux.POST("/cart/create", mux.CORS(m.Auth(www.Cart.Create)))
	mux.GET("/cart/:id", m.Auth(www.Cart.Get))
	mux.GET("/carts", m.Auth(www.Cart.GetAll))
	mux.PUT("/cart/update/:id", mux.CORS(m.Auth(www.Cart.Update)))
	mux.DEL("/cart/delete/:id", mux.CORS(m.Auth(www.Cart.Delete)))

	mux.ServeStaticFiles("static")

	return mux
}
