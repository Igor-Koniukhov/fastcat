package handlers

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/controllers"
	"net/http"
)

func CartHandlers(db *sql.DB,)  {
	controller := controllers.CartControllers{}

	http.HandleFunc("/supplier/create", controller.CreateCart(db))
	http.HandleFunc("/supplier/", controller.GetCart(db))
	http.HandleFunc("/supplier", controller.GetAllCarts(db))
	http.HandleFunc("/supplier/update/", controller.UpdateCart(db))
	http.HandleFunc("/supplier/delete/", controller.DeleteCart(db))
}