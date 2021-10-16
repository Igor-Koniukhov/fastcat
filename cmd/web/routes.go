package main

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/handlers"
	"github.com/igor-koniukhov/fastcat/internal/repository"
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
	mux.GET("/status", www.Auth(www.User.StatusPage))
	mux.GET("/contacts", www.User.Contacts)
	mux.GET("/user/:id", mux.CORS(www.Auth(www.User.Get)))
	mux.GET("/users", www.User.GetAll)
	mux.PUT("/user/update/:id", mux.CORS(www.Auth(www.User.Update)))
	mux.DEL("/user/delete/:id", mux.CORS(www.Auth(www.User.Delete)))

	mux.GET("/order/blank", www.Auth(www.Order.ShowBlankOrder))
	mux.POST("/order/create", mux.CORS(www.Auth(www.Order.Create)))
	mux.GET("/order/:id", mux.CORS(www.Auth(www.Order.Get)))
	mux.GET("/orders", mux.CORS(www.Auth(www.Order.GetAll)))
	mux.PUT("/order/update/:id", mux.CORS(www.Auth(www.Order.Update)))

	mux.GET("/supplier/:id", www.Auth(www.Supplier.Get))
	mux.GET("/suppliers", www.Auth(www.Supplier.GetAll))
	mux.GET("/suppliers/selected", www.Auth(www.Supplier.GetAllBySchedule))
	mux.PUT("/supplier/update/:id", mux.CORS(www.Auth(www.Supplier.Update)))
	mux.DEL("/supplier/delete/:id", mux.CORS(www.Auth(www.Supplier.Delete)))

	mux.GET("/product/:id", www.Auth(www.Product.Get))
	mux.GET("/suppliers-products/", www.Auth(www.Product.GetAllBySupplierID))
	mux.GET("/products", www.Auth(www.Product.GetAll))
	mux.GET("/products-json", mux.JSON(www.Product.GetJson))
	mux.PUT("/product/update/:id", mux.CORS(www.Auth(www.Product.Update)))
	mux.DEL("/product/delete/:id", mux.CORS(www.Auth(www.Product.Delete)))

	mux.POST("/cart/create", mux.CORS(www.Auth(www.Cart.Create)))
	mux.GET("/cart/:id", www.Auth(www.Cart.Get))
	mux.GET("/carts", www.Auth(www.Cart.GetAll))
	mux.PUT("/cart/update/:id", mux.CORS(www.Auth(www.Cart.Update)))
	mux.DEL("/cart/delete/:id", mux.CORS(www.Auth(www.Cart.Delete)))

	mux.ServeStaticFiles("static")

	return mux
}
