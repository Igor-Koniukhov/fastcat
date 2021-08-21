package handlers

import (
	"github.com/igor-koniukhov/fastcat/controllers"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"net/http"
)

func UserHandlers(app *config.AppConfig)  {
	repo := controllers.NewUserControllers(app)
	controllers.NewControllers(repo)
	app.Str = "hello from main"

	http.HandleFunc("/user/create", controllers.Repo.CreateUser( "POST"))
	// GetUser able search by id and email
	http.HandleFunc("/user/", controllers.Repo.GetUser( "GET"))
	http.HandleFunc("/users", controllers.Repo.GetAllUsers( "GET"))
	http.HandleFunc("/user/update/", controllers.Repo.UpdateUser( "PUT"))
	http.HandleFunc("/user/delete/", controllers.Repo.DeleteUser( "DELETE"))

}