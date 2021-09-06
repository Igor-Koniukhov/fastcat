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

	http.HandleFunc("/cart/create", cr.CreateCart("POST"))
	http.HandleFunc("/cart/", cr.GetCart("GET"))
	http.HandleFunc("/cart", cr.GetAllCarts("GET"))
	http.HandleFunc("/cart/update/", cr.UpdateCart("PUT"))
	http.HandleFunc("/cart/delete/", cr.DeleteCart("DELETE"))
}
