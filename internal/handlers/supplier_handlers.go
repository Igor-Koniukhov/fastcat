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

	http.HandleFunc("/supplier/create", sr.Create("POST"))
	http.HandleFunc("/supplier/", sr.Get("GET"))
	http.HandleFunc("/suppliers", sr.GetAll("GET"))
	http.HandleFunc("/supplier/update/", sr.Update("PUT"))
	http.HandleFunc("/supplier/delete/", sr.Delete("DELETE"))
}
