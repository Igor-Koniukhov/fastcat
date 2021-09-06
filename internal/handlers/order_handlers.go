package handlers

import (

	"github.com/igor-koniukhov/fastcat/controllers"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"net/http"
)

func OrderHandlers( app *config.AppConfig) {

	repo := controllers.NewOrderControllers(app)
	controllers.NewControllersO(repo)
	or := controllers.RepoOrder

	http.HandleFunc("/order/create", or.Create("POST"))
	http.HandleFunc("/order/", or.Get("GET"))
	http.HandleFunc("/orders", or.GetAll("GET"))
	http.HandleFunc("/order/update/", or.Update("PUT"))
	http.HandleFunc("/order/delete/", or.Delete("DELETE"))

}
