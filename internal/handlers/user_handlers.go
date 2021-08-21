package handlers

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/controllers"
	"net/http"
)

func UserHandlers(db *sql.DB)  {
	controller := controllers.UserControllers{}

	http.HandleFunc("/user/create", controller.CreateUser(db))
	// GetUser able search by id and email
	http.HandleFunc("/user/", controller.GetUser(db))
	http.HandleFunc("/users", controller.GetAllUsers(db))
	http.HandleFunc("/user/update/", controller.UpdateUser(db))
	http.HandleFunc("/user/delete/", controller.DeleteUser(db))
}