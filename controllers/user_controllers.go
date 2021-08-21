package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"log"
	"net/http"

)

type ControllerInterface interface {
	CreateUser(db *sql.DB) http.HandlerFunc
	GetUser(db *sql.DB) http.HandlerFunc
	GetAllUsers(db *sql.DB) http.HandlerFunc
	DeleteUser(db *sql.DB) http.HandlerFunc
	UpdateUser(db *sql.DB) http.HandlerFunc
}

type Controllers struct {
	App *config.AppConfig
}

const TableUser = "user"

var user model.User

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controllers) CreateUser(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var u model.User
		switch r.Method {
		case "POST":
			json.NewDecoder(r.Body).Decode(&u)
			userRepo := repository.UserRepository{}
			user, err := userRepo.CreateUser(&u, db)
			checkError(err)
			json.NewEncoder(w).Encode(&user)

		default:
			methodMassage(w, "POST")
		}
	}
}

var email = "Jone@gmail.com"

func (c Controllers) GetUser( db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			userRepo := repository.UserRepository{}
			param, nameParam, _ := userRepo.Param(r)
			fmt.Println(param, nameParam)
			user := userRepo.GetUser(&nameParam, &param, db)
			fmt.Fprintf(w, user.Name)
		default:
			methodMassage(w, "GET")
		}
	}
}

func (c Controllers) GetAllUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			userRepo := repository.UserRepository{}
			users := userRepo.GetAllUsers(db)
			json.NewEncoder(w).Encode(&users)
		default:
			methodMassage(w, "GET")
		}
	}
}

func (c Controllers) DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "DELETE":

			userRepo := repository.UserRepository{}
			_, _, id := userRepo.Param(r)
			err := userRepo.DeleteUser(id, db)
			checkError(err)
			_, _ = fmt.Fprintf(w, fmt.Sprintf(" user with %d deleted", id))
		default:
			methodMassage(w, "DELETE")
		}
	}
}

func (c Controllers) UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			var id = 5
			var u model.User
			json.NewDecoder(r.Body).Decode(&u)
			userRepo := repository.UserRepository{}
			user := userRepo.UpdateUser(id, &u, db)
			json.NewEncoder(w).Encode(&user)
		default:
			methodMassage(w, "UPDATE")
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
