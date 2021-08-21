package handlers

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/controllers"
	"net/http"
)

func ProductHandlers(db *sql.DB,)  {
	controller := controllers.ProductControllers{}

	http.HandleFunc("/product/create", controller.CreateProduct(db))
	http.HandleFunc("/product/", controller.GetProduct(db))
	http.HandleFunc("/product", controller.GetAllProducts(db))
	http.HandleFunc("/product/update/", controller.UpdateProduct(db))
	http.HandleFunc("/product/delete/", controller.DeleteProduct(db))
}