package controllers

import (

	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"log"
	"net/http"

)

type UserControllerI interface {
	CreateUser( method string) http.HandlerFunc
	GetUser( method string) http.HandlerFunc
	GetAllUsers( method string) http.HandlerFunc
	DeleteUser( method string) http.HandlerFunc
	UpdateUser( method string) http.HandlerFunc
}

var Repo *UserControllers

type UserControllers struct {
	App *config.AppConfig
}

func NewUserControllers(app *config.AppConfig) *UserControllers {
	return &UserControllers{App: app}
}

func NewControllers(r *UserControllers)  {
	Repo = r
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c *UserControllers) CreateUser( method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u model.User
		switch r.Method {
		case method:
			json.NewDecoder(r.Body).Decode(&u)
			userRepo := repository.UserRepository{}
			user, err := userRepo.CreateUser(&u, c.App.DB)
			checkError(err)
			json.NewEncoder(w).Encode(&user)

		default:
			methodMassage(w, method)
		}
	}
}



func (c *UserControllers) GetUser(  method string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(c.App.Session)
		switch r.Method {
		case method:
			userRepo := repository.UserRepository{}
			param, nameParam, _ := userRepo.Param(r)
			fmt.Println(param, nameParam)
			user := userRepo.GetUser(&nameParam, &param, c.App.DB)
			fmt.Fprintf(w, user.Name)
		default:
			methodMassage(w, method)
		}
	}
}

func (c *UserControllers) GetAllUsers( method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			userRepo := repository.UserRepository{}
			users := userRepo.GetAllUsers(c.App.DB)
			json.NewEncoder(w).Encode(&users)
		default:
			methodMassage(w, method)
		}
	}
}

func (c *UserControllers) DeleteUser( method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:

			userRepo := repository.UserRepository{}
			_, _, id := userRepo.Param(r)
			err := userRepo.DeleteUser(id, c.App.DB)
			checkError(err)
			_, _ = fmt.Fprintf(w, fmt.Sprintf(" user with %d deleted", id))
		default:
			methodMassage(w, method)
		}
	}
}

func (c *UserControllers) UpdateUser( method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			var id = 5
			var u model.User
			json.NewDecoder(r.Body).Decode(&u)
			userRepo := repository.UserRepository{}
			user := userRepo.UpdateUser(id, &u, c.App.DB)
			json.NewEncoder(w).Encode(&user)
		default:
			methodMassage(w, method)
		}
	}
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
func methodMassage(w http.ResponseWriter, m string) {
	http.Error(w, "Only "+m+" method is allowed", http.StatusMethodNotAllowed)

}
