package handlers

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/controllers"
	"net/http"
)

func SupplierHandlers(db *sql.DB,)  {
	controller := controllers.SupplierControllers{}

	http.HandleFunc("/supplier/create", controller.CreateSupplier(db))
	http.HandleFunc("/supplier/", controller.GetSupplier(db))
	http.HandleFunc("/supplier", controller.GetAllSuppliers(db))
	http.HandleFunc("/supplier/update/", controller.UpdateSupplier(db))
	http.HandleFunc("/supplier/delete/", controller.DeleteSupplier(db))
}