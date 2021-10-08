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
	mux.GET("/", www.Supplier.Home)
	mux.GET("/about", www.User.AboutUs)
	mux.GET("/contacts", www.User.Contacts)
	mux.POST("/login", mux.CORS(www.User.PostLogin))
	mux.GET("/show-login", www.User.ShowLogin)
	mux.GET("/refresh", mux.CORS(m.Auth(www.User.Refresh)))
	mux.POST("/logout", mux.CORS(www.User.Logout))

	mux.GET("/registration", www.User.ShowRegistration)
	mux.POST("/signup", mux.CORS(www.User.SingUp))
	mux.GET("/user/:id", mux.CORS(m.Auth(www.User.Get)))
	mux.GET("/users", www.User.GetAll)
	mux.PUT("/user/update/:id", mux.CORS(m.Auth(www.User.Update)))
	mux.DEL("/user/delete/:id", mux.CORS(m.Auth(www.User.Delete)))

	mux.GET("/order/blank", www.Order.ShowBlankOrder)
	mux.POST("/order/create", mux.CORS(www.Order.Create))
	mux.GET("/order/:id", mux.CORS(m.Auth(www.Order.Get)))
	mux.GET("/orders", mux.CORS(m.Auth(www.Order.GetAll)))
	mux.PUT("/order/update/:id", mux.CORS(m.Auth(www.Order.Update)))

	mux.GET("/supplier/:id", www.Supplier.Get)
	mux.GET("/suppliers", www.Supplier.GetAll)
	mux.GET("/suppliers/selected", www.Supplier.GetAllBySchedule)
	mux.PUT("/supplier/update/:id", mux.CORS(m.Auth(www.Supplier.Update)))
	mux.DEL("/supplier/delete/:id", mux.CORS(m.Auth(www.Supplier.Delete)))

	mux.GET("/product/:id", www.Product.Get)
	mux.GET("/suppliers-products", www.Product.GetAllBySupplierID)
	mux.GET("/products", www.Product.GetAll)
	mux.GET("/products-json", mux.JSON(www.Product.GetJson))
	mux.GET("/fetch", mux.CORS(www.Product.FetchAll))
	mux.PUT("/product/update/:id", mux.CORS(m.Auth(www.Product.Update)))
	mux.DEL("/product/delete/:id", mux.CORS(m.Auth(www.Product.Delete)))

	mux.POST("/cart/create", mux.CORS(www.Cart.Create))
	mux.GET("/cart/:id", www.Cart.Get)
	mux.GET("/carts", www.Cart.GetAll)
	mux.PUT("/cart/update/:id", mux.CORS(m.Auth(www.Cart.Update)))
	mux.DEL("/cart/delete/:id", mux.CORS(m.Auth(www.Cart.Delete)))

	mux.ServeStaticFiles("static")

	return mux
}
