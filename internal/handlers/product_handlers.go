package handlers

import (
	"github.com/igor-koniukhov/fastcat/controllers"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"net/http"
)

func ProductHandlers(app *config.AppConfig)  {

	repo := controllers.NewProductControllers(app)
	controllers.NewControllersP(repo)
	pc := controllers.RepoProducts

	http.HandleFunc("/product/create", pc.Update( "POST"))
	http.HandleFunc("/product/", pc.Get( "GET"))
	http.HandleFunc("/products", pc.GetAll( "GET"))
	http.HandleFunc("/product/update/", pc.Update( "PUT"))
	http.HandleFunc("/product/delete/", pc.Delete( "DELETE"))
}