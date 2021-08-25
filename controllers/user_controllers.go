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
	CreateUser(method string) http.HandlerFunc
	GetUser(method string) http.HandlerFunc
	GetAllUsers(method string) http.HandlerFunc
	DeleteUser(method string) http.HandlerFunc
	UpdateUser(method string) http.HandlerFunc
}

var RepoUser *UserControllers

type UserControllers struct {
	App *config.AppConfig
}

func NewUserControllers(app *config.AppConfig) *UserControllers {
	return &UserControllers{App: app}
}

func NewControllersU(r *UserControllers) {
	RepoUser = r
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func UserAppConfigProvider(a *config.AppConfig) *repository.UserRepository {
	repo := repository.NewUserRepository(a)
	repository.NewRepoU(repo)
	return repo

}

func (c *UserControllers) CreateUser(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u model.User
		switch r.Method {
		case method:
			json.NewDecoder(r.Body).Decode(&u)
			UserAppConfigProvider(c.App)
			user, err := repository.RepoU.CreateUser(&u)
			checkError(err)
			json.NewEncoder(w).Encode(&user)

		default:
			methodMessage(w, method)
		}
	}
}

func (c *UserControllers) GetUser(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			UserAppConfigProvider(c.App)
			param, nameParam, _ := repository.RepoU.Param(r)
			user := repository.RepoU.GetUser(&nameParam, &param)
			json.NewEncoder(w).Encode(&user)
		default:
			methodMessage(w, method)
		}
	}
}

func (c *UserControllers) GetAllUsers(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			UserAppConfigProvider(c.App)
			users := repository.RepoU.GetAllUsers()
			json.NewEncoder(w).Encode(&users)
		default:
			methodMessage(w, method)
		}
	}
}

func (c *UserControllers) DeleteUser(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:

			repo := UserAppConfigProvider(c.App)
			_, _, id := repo.Param(r)
			err := repository.RepoU.DeleteUser(id)
			checkError(err)
			_, _ = fmt.Fprintf(w, fmt.Sprintf(" user with %d deleted", id))
		default:
			methodMessage(w, method)
		}
	}
}

func (c *UserControllers) UpdateUser(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			var u model.User
			json.NewDecoder(r.Body).Decode(&u)
			repo := UserAppConfigProvider(c.App)
			_, _, id := repo.Param(r)
			user := repository.RepoU.UpdateUser(id, &u)
			json.NewEncoder(w).Encode(&user)
		default:
			methodMessage(w, method)
		}
	}
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
func methodMessage(w http.ResponseWriter, m string) {
	http.Error(w, "Only "+m+" method is allowed", http.StatusMethodNotAllowed)

}
