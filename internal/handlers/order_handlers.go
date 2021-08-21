package handlers

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/controllers"
	"net/http"
)

func OrderHandlers(db *sql.DB,)  {
	controller := controllers.OrderControllers{}

	http.HandleFunc("/supplier/create", controller.CreateOrder(db))
	http.HandleFunc("/supplier/", controller.GetOrder(db))
	http.HandleFunc("/supplier", controller.GetAllOrders(db))
	http.HandleFunc("/supplier/update/", controller.UpdateOrder(db))
	http.HandleFunc("/supplier/delete/", controller.DeleteOrder(db))
}