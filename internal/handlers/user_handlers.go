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

	http.HandleFunc("/user/create", ur.Create( "POST"))
	http.HandleFunc("/user/", ur.Get("GET"))
	http.HandleFunc("/users", ur.GetAllU( "GET"))
	http.HandleFunc("/user/update/", ur.Update( "PUT"))
	http.HandleFunc("/user/delete/", ur.Delete( "DELETE"))

}