package handlers

import (
	"github.com/igor-koniukhov/fastcat/controllers"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"net/http"
)

func UserHandlers(app *config.AppConfig)  {

	repo := controllers.NewUserControllers(app)
	controllers.NewControllersU(repo)
	ur := controllers.RepoUser

	http.HandleFunc("/user/create", ur.CreateUser( "POST"))
	// GetUser able search by id and email
	http.HandleFunc("/user/", ur.GetUser( "GET"))
	http.HandleFunc("/users", ur.GetAllUsers( "GET"))
	http.HandleFunc("/user/update/", ur.UpdateUser( "PUT"))
	http.HandleFunc("/user/delete/", ur.DeleteUser( "DELETE"))

}