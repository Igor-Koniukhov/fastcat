package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/igor-koniukhov/fastcat/controllers"
	"github.com/igor-koniukhov/fastcat/driver"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
)

var app config.AppConfig

func init() {
	gotenv.Load()
}

func main() {
	db := driver.ConnectMySQLDB()
	defer db.Close()
	port := os.Getenv("PORT")

	SetAppConfigParameters(db)
	SetWebLoggerParameters()
	go RunUpToDateSuppliersInfo(600)

	repo := repository.NewRepository(&app)
	c := controllers.NewControllers(repo)
	controllers.NewController(c)
	repository.NewRepo(repo)

	http.HandleFunc("/user/create", c.User.Create())
	http.HandleFunc("/user/", c.User.Get())
	http.HandleFunc("/users", c.User.GetAll())
	http.HandleFunc("/user/update/", c.User.Update())
	http.HandleFunc("/user/delete/", c.User.Delete())

	http.HandleFunc("/order/create", c.Order.Create())
	http.HandleFunc("/order/", c.Order.Get())
	http.HandleFunc("/orders", c.Order.GetAll())
	http.HandleFunc("/order/update/", c.Order.Update())
	http.HandleFunc("/order/delete/", c.Order.Delete())

	http.HandleFunc("/supplier/create", c.Supplier.Create())
	http.HandleFunc("/supplier/", c.Supplier.Get())
	http.HandleFunc("/suppliers", c.Supplier.GetAll())
	http.HandleFunc("/supplier/update/", c.Supplier.Update())
	http.HandleFunc("/supplier/delete/", c.Supplier.Delete())

	http.HandleFunc("/product/create", c.Product.Update())
	http.HandleFunc("/product/", c.Product.Get())
	http.HandleFunc("/products", c.Product.GetAll())
	http.HandleFunc("/product/update/", c.Product.Update())
	http.HandleFunc("/product/delete/", c.Product.Delete())

	http.HandleFunc("/cart/create", c.Cart.Create())
	http.HandleFunc("/cart/", c.Cart.Get())
	http.HandleFunc("/cart", c.Cart.GetAll())
	http.HandleFunc("/cart/update/", c.Cart.Update())
	http.HandleFunc("/cart/delete/", c.Cart.Delete())

	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Fatal(http.ListenAndServe(port, nil))
}
