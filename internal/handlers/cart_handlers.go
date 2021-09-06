package handlers

import (
	"github.com/igor-koniukhov/fastcat/controllers"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"net/http"
)

func CartHandlers(app *config.AppConfig) {

	repo := controllers.NewCartControllers(app)
	controllers.NewControllersC(repo)
	cr := controllers.RepoCart

	http.HandleFunc("/cart/create", cr.Create("POST"))
	http.HandleFunc("/cart/", cr.Get("GET"))
	http.HandleFunc("/cart", cr.GetAll("GET"))
	http.HandleFunc("/cart/update/", cr.Update("PUT"))
	http.HandleFunc("/cart/delete/", cr.Delete("DELETE"))
}
