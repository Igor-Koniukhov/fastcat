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
	http.HandleFunc("/login", ur.Login("POST"))
	http.HandleFunc("/refresh", ur.Refresh("POST"))
	http.HandleFunc("/profile", ur.GetProfile("GET"))
	http.HandleFunc("/users", ur.GetAllUsers( "GET"))
	http.HandleFunc("/user/update/", ur.UpdateUser( "PUT"))
	http.HandleFunc("/user/delete/", ur.DeleteUser( "DELETE"))

}