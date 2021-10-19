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

	http.HandleFunc("/order/create", or.CreateOrder("POST"))
	http.HandleFunc("/order/", or.GetOrder("GET"))
	http.HandleFunc("/orders", or.GetAllOrders("GET"))
	http.HandleFunc("/order/update/", or.UpdateOrder("PUT"))
	http.HandleFunc("/order/delete/", or.DeleteOrder("DELETE"))

}
