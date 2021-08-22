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

	http.HandleFunc("/product/create", pc.CreateProduct( "POST"))
	http.HandleFunc("/product/", pc.GetProduct( "GET"))
	http.HandleFunc("/product", pc.GetAllProducts( "GET"))
	http.HandleFunc("/product/update/", pc.UpdateProduct( "PUT"))
	http.HandleFunc("/product/delete/", pc.DeleteProduct( "DELETE"))
}