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
	Create(method string) http.HandlerFunc
	Get(method string) http.HandlerFunc
	GetAllU(method string) http.HandlerFunc
	Delete(method string) http.HandlerFunc
	Update(method string) http.HandlerFunc
}

var RepoUser *UserController

type UserController struct {
	App *config.AppConfig
}

func NewUserControllers(app *config.AppConfig) *UserController {
	return &UserController{App: app}
}

func NewControllersU(r *UserController) {
	RepoUser = r
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func userAppConfigProvider(a *config.AppConfig) *repository.UserRepository {
	repo := repository.NewUserRepository(a)
	repository.NewRepoU(repo)
	return repo
}

func (c *UserController) Create(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u model.User
		switch r.Method {
		case method:
			json.NewDecoder(r.Body).Decode(&u)
			_ = userAppConfigProvider(c.App)
			user, err := repository.RepoU.Create(&u)
			checkError(err)
			json.NewEncoder(w).Encode(&user)
		default:
			methodMessage(w, method)
		}
	}
}

func (c *UserController) Get(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		switch r.Method {
		case method:
			repo := userAppConfigProvider(c.App)
			param, nameParam, _ := repo.Param(r)
			user := repository.RepoU.Get(&nameParam, &param)
			json.NewEncoder(w).Encode(&user)
		default:
			methodMessage(w, method)
		}
	}
}

func (c *UserController) GetAllU(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		switch r.Method {
		case method:
			_ = userAppConfigProvider(c.App)
			users := repository.RepoU.GetAll()
			json.NewEncoder(w).Encode(&users)
		default:
			methodMessage(w, method)
		}
	}
}

func (c *UserController) Delete(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			repo := userAppConfigProvider(c.App)
			_, _, id := repo.Param(r)
			err := repository.RepoU.Delete(id)
			checkError(err)
			_, _ = fmt.Fprintf(w, fmt.Sprintf(" user with %d deleted", id))
		default:
			methodMessage(w, method)
		}
	}
}

func (c *UserController) Update(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			var u model.User
			json.NewDecoder(r.Body).Decode(&u)
			repo := userAppConfigProvider(c.App)
			_, _, id := repo.Param(r)
			user := repository.RepoU.Update(id, &u)
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
