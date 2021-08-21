package handlers

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/controllers"
	"net/http"
)

func CartHandlers(db *sql.DB,)  {
	controller := controllers.CartControllers{}

	http.HandleFunc("/cart/create", controller.CreateCart(db, "POST"))
	http.HandleFunc("/cart/", controller.GetCart(db, "GET"))
	http.HandleFunc("/cart", controller.GetAllCarts(db, "GET"))
	http.HandleFunc("/cart/update/", controller.UpdateCart(db, "PUT"))
	http.HandleFunc("/cart/delete/", controller.DeleteCart(db, "DELETE"))
}