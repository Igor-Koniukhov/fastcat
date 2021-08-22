package handlers

import (
	"github.com/igor-koniukhov/fastcat/controllers"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"net/http"
)

func SupplierHandlers(app *config.AppConfig) {

	repo := controllers.NewSupplierControllers(app)
	controllers.NewControllersS(repo)
	sr := controllers.RepoSupplier

	http.HandleFunc("/supplier/create", sr.CreateSupplier("POST"))
	http.HandleFunc("/supplier/", sr.GetSupplier("GET"))
	http.HandleFunc("/supplier", sr.GetAllSuppliers("GET"))
	http.HandleFunc("/supplier/update/", sr.UpdateSupplier("PUT"))
	http.HandleFunc("/supplier/delete/", sr.DeleteSupplier("DELETE"))
}
