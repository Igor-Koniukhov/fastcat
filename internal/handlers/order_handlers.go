package handlers

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/controllers"
	"net/http"
)

func OrderHandlers(db *sql.DB,)  {

	controller := controllers.OrderControllers{}

	http.HandleFunc("/order/create", controller.CreateOrder(db, "POST"))
	http.HandleFunc("/order/", controller.GetOrder(db, "GET"))
	http.HandleFunc("/order", controller.GetAllOrders(db, "GET"))
	http.HandleFunc("/order/update/", controller.UpdateOrder(db, "PUT"))
	http.HandleFunc("/order/delete/", controller.DeleteOrder(db, "DELETE"))
}