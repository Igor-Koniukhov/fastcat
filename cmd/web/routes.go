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

	mux.GET("/", mux.CORS(www.AuthSet(www.Supplier.Home)))
	mux.GET("/registration", mux.CORS(www.AuthSet(www.User.ShowRegistration)))
	mux.POST("/sign-up", mux.CORS(www.AuthSet(www.User.SingUp)))
	mux.GET("/show-login", www.AuthSet(www.User.ShowLogin))
	mux.POST("/login", mux.CORS(www.AuthSet(www.User.PostLogin)))
	mux.GET("/refresh", www.AuthSet(www.User.RefreshToken))
	mux.GET("/logout", www.AuthSet(www.User.Logout))

	mux.GET("/about", www.AuthSet(www.User.AboutUs))
	mux.GET("/contacts", www.AuthSet(www.User.Contacts))
	mux.GET("/user/:id", mux.CORS(www.AuthSet(www.AuthCheck(www.User.Get))))
	mux.GET("/users", www.AuthSet(www.User.GetAll))
	mux.PUT("/user/update/:id", mux.CORS(www.AuthSet(www.AuthCheck(www.User.Update))))
	mux.DEL("/user/delete/:id", mux.CORS(www.AuthSet(www.AuthCheck(www.User.Delete))))

	mux.GET("/order/blank", www.AuthSet(www.AuthCheck(www.Order.ShowBlankOrder)))
	mux.POST("/order/create", mux.CORS(www.AuthSet(www.AuthCheck(www.Order.Create))))
	mux.GET("/order/:id", mux.CORS(www.AuthSet(www.AuthCheck(www.Order.Get))))
	mux.GET("/orders", mux.CORS(www.AuthSet(www.AuthCheck(www.Order.GetAll))))
	mux.PUT("/order/update/:id", mux.CORS(www.AuthSet(www.AuthCheck(www.Order.Update))))

	mux.GET("/supplier/:id", www.AuthSet(www.AuthCheck(www.Supplier.Get)))
	mux.GET("/suppliers",www.AuthSet(www.AuthCheck(www.Supplier.GetAll)))
	mux.GET("/suppliers/type", www.AuthSet(www.AuthCheck(www.Supplier.GetAllByType)))
	mux.GET("/suppliers/selected", www.AuthSet(www.AuthCheck(www.Supplier.GetAllBySchedule)))
	mux.PUT("/supplier/update/:id", mux.CORS(www.AuthSet(www.AuthCheck(www.Supplier.Update))))
	mux.DEL("/supplier/delete/:id", mux.CORS(www.AuthSet(www.AuthCheck(www.Supplier.Delete))))

	mux.GET("/product/:id", www.AuthSet(www.AuthCheck(www.Product.Get)))
	mux.GET("/suppliers-products", www.AuthSet(www.AuthCheck(www.Product.GetAllBySupplierID)))
	mux.GET("/products/type", www.AuthSet(www.AuthCheck(www.Product.GetAllByType)))
	mux.GET("/products", www.AuthSet(www.AuthCheck(www.Product.GetAll)))
	mux.GET("/products-json", mux.JSON(www.Product.GetJson))
	mux.PUT("/product/update/:id", mux.CORS(www.AuthSet(www.AuthCheck(www.Product.Update))))
	mux.DEL("/product/delete/:id", mux.CORS(www.AuthSet(www.AuthCheck(www.Product.Delete))))

	mux.POST("/cart/create", mux.CORS(www.AuthSet(www.AuthCheck(www.Cart.Create))))
	mux.GET("/cart/:id", www.AuthSet(www.AuthCheck(www.Cart.Get)))
	mux.GET("/carts", www.AuthSet(www.AuthCheck(www.Cart.GetAll)))
	mux.GET("/cabinet", mux.CORS(www.AuthSet(www.AuthCheck(www.Cart.GetAllByUserID))))
	mux.PUT("/cart/update/:id", mux.CORS(www.AuthSet(www.AuthCheck(www.Cart.Update))))
	mux.DEL("/cart/delete/:id", mux.CORS(www.AuthSet(www.AuthCheck(www.Cart.Delete))))

	mux.ServeStaticFiles("static")

	return mux
}
