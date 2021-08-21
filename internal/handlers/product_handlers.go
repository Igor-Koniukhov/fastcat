package handlers

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/controllers"
	"net/http"
)

func ProductHandlers(db *sql.DB,)  {

	controller := controllers.ProductControllers{}

	http.HandleFunc("/product/create", controller.CreateProduct(db, "POST"))
	http.HandleFunc("/product/", controller.GetProduct(db, "GET"))
	http.HandleFunc("/product", controller.GetAllProducts(db, "GET"))
	http.HandleFunc("/product/update/", controller.UpdateProduct(db, "PUT"))
	http.HandleFunc("/product/delete/", controller.DeleteProduct(db, "DELETE"))
}