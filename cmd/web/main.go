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
	go RunUpToDateSuppliersInfo(60)

	repo := repository.NewRepository(&app)
	c := controllers.NewControllers(repo)
	controllers.NewController(c)
	repository.NewRepo(repo)

	http.HandleFunc("/user/create", c.User.Create( "POST"))
	http.HandleFunc("/user/", c.User.Get("GET"))
	http.HandleFunc("/users", c.User.GetAll( "GET"))
	http.HandleFunc("/user/update/", c.User.Update( "PUT"))
	http.HandleFunc("/user/delete/", c.User.Delete( "DELETE"))

	http.HandleFunc("/order/create", c.Order.Create("POST"))
	http.HandleFunc("/order/", c.Order.Get("GET"))
	http.HandleFunc("/orders", c.Order.GetAll("GET"))
	http.HandleFunc("/order/update/", c.Order.Update("PUT"))
	http.HandleFunc("/order/delete/", c.Order.Delete("DELETE"))

	http.HandleFunc("/supplier/create", c.Supplier.Create("POST"))
	http.HandleFunc("/supplier/", c.Supplier.Get("GET"))
	http.HandleFunc("/suppliers", c.Supplier.GetAll("GET"))
	http.HandleFunc("/supplier/update/", c.Supplier.Update("PUT"))
	http.HandleFunc("/supplier/delete/", c.Supplier.Delete("DELETE"))

	http.HandleFunc("/product/create", c.Product.Update( "POST"))
	http.HandleFunc("/product/", c.Product.Get( "GET"))
	http.HandleFunc("/products", c.Product.GetAll( "GET"))
	http.HandleFunc("/product/update/", c.Product.Update( "PUT"))
	http.HandleFunc("/product/delete/", c.Product.Delete( "DELETE"))

	http.HandleFunc("/cart/create", c.Cart.Create("POST"))
	http.HandleFunc("/cart/", c.Cart.Get("GET"))
	http.HandleFunc("/cart", c.Cart.GetAll("GET"))
	http.HandleFunc("/cart/update/", c.Cart.Update("PUT"))
	http.HandleFunc("/cart/delete/", c.Cart.Delete("DELETE"))



	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Fatal(http.ListenAndServe(port, nil))
}





