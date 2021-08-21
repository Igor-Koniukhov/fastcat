package handlers

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/controllers"
	"net/http"
)

func SupplierHandlers(db *sql.DB,)  {

	controller := controllers.SupplierControllers{}

	http.HandleFunc("/supplier/create", controller.CreateSupplier(db, "POST"))
	http.HandleFunc("/supplier/", controller.GetSupplier(db, "GET"))
	http.HandleFunc("/supplier", controller.GetAllSuppliers(db, "GET"))
	http.HandleFunc("/supplier/update/", controller.UpdateSupplier(db, "PUT"))
	http.HandleFunc("/supplier/delete/", controller.DeleteSupplier(db, "DELETE"))
}