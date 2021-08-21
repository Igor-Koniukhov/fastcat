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

type UserControllerI interface {
	CreateUser(db *sql.DB, method string) http.HandlerFunc
	GetUser(db *sql.DB, method string) http.HandlerFunc
	GetAllUsers(db *sql.DB, method string) http.HandlerFunc
	DeleteUser(db *sql.DB, method string) http.HandlerFunc
	UpdateUser(db *sql.DB, method string) http.HandlerFunc
}

type UserControllers struct {
	App *config.AppConfig
}



var user model.User

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c UserControllers) CreateUser(db *sql.DB, method string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var u model.User
		switch r.Method {
		case method:
			json.NewDecoder(r.Body).Decode(&u)
			userRepo := repository.UserRepository{}
			user, err := userRepo.CreateUser(&u, db)
			checkError(err)
			json.NewEncoder(w).Encode(&user)

		default:
			methodMassage(w, method)
		}
	}
}



func (c UserControllers) GetUser( db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			userRepo := repository.UserRepository{}
			param, nameParam, _ := userRepo.Param(r)
			fmt.Println(param, nameParam)
			user := userRepo.GetUser(&nameParam, &param, db)
			fmt.Fprintf(w, user.Name)
		default:
			methodMassage(w, method)
		}
	}
}

func (c UserControllers) GetAllUsers(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			userRepo := repository.UserRepository{}
			users := userRepo.GetAllUsers(db)
			json.NewEncoder(w).Encode(&users)
		default:
			methodMassage(w, method)
		}
	}
}

func (c UserControllers) DeleteUser(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:

			userRepo := repository.UserRepository{}
			_, _, id := userRepo.Param(r)
			err := userRepo.DeleteUser(id, db)
			checkError(err)
			_, _ = fmt.Fprintf(w, fmt.Sprintf(" user with %d deleted", id))
		default:
			methodMassage(w, method)
		}
	}
}

func (c UserControllers) UpdateUser(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			var id = 5
			var u model.User
			json.NewDecoder(r.Body).Decode(&u)
			userRepo := repository.UserRepository{}
			user := userRepo.UpdateUser(id, &u, db)
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
