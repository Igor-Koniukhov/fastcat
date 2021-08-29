package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/auth/services"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"log"
	"net/http"
	"strings"
)

type UserControllerI interface {
	CreateUser(method string) http.HandlerFunc
	Login(method string) http.HandlerFunc
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

func userAppConfigProvider(a *config.AppConfig) *repository.UserRepository {
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
			_ = userAppConfigProvider(c.App)
			user, err := repository.RepoU.CreateUser(&u)
			c.App.ErrorLog.Println(err, "error from appConfig")
			checkError(err)
			json.NewEncoder(w).Encode(&user)

		default:
			methodMessage(w, method)
		}
	}
}

func (c *UserControllers) Login(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			logReq := new(model.LoginRequest)
			if err := json.NewDecoder(r.Body).Decode(&logReq); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			resp, err := services.TokenResponder(w, logReq, c.App)
			if err != nil {
				log.Println(err)
			}

			if err != nil {
				c.App.ErrorLog.Println(err)
			}
			c.App.AccessToken = resp.AccessToken
			c.App.RefreshToken = resp.RefreshToken

			str := c.App.RefreshToken
			partAccess := strings.Split(str, ".")
			thirdPartAccess := partAccess[len(partAccess)-1]
			w.Header().Set("Authorization", c.App.AccessToken)
			err = json.NewEncoder(w).Encode(resp)

			fmt.Println(partAccess)
			fmt.Println(thirdPartAccess)

		default:
			methodMessage(w, method)
		}
	}
}

func (c *UserControllers) Refresh(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			logReq := new(model.LoginRequest)
			if err := json.NewDecoder(r.Body).Decode(&logReq); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			resp, err := services.TokenResponder(w, logReq, c.App)
			if err != nil {
				log.Println(err)
			}

			w.WriteHeader(http.StatusOK)
			err = json.NewEncoder(w).Encode(resp)
			if err != nil {
				log.Println(err)
			}

		default:
			methodMessage(w, method)
		}
	}

}
func (c *UserControllers) GetProfile(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case method:
			bearerString := r.Header.Get("Authorization")
			fmt.Println(bearerString, "this is bearer string")
			tokenString := services.GetTokenFromBearerString(bearerString)
			claims, err := services.ValidateToken(tokenString, services.AccessSecret)
			//claims, err := ValidateToken(GetTokenFromBearerString(r.Header.Get("Authorization")), RefreshSecret)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			userAppConfigProvider(c.App)
			user, err := repository.RepoU.GetUserByID(claims.ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			resp := &services.UserResponse{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
			}

			w.WriteHeader(http.StatusOK)
			err = json.NewEncoder(w).Encode(resp)
			if err != nil {
				log.Println(err)
			}

		default:
			methodMessage(w, method)

		}
	}
}

func (c *UserControllers) GetAllUsers(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			_ = userAppConfigProvider(c.App)
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
			repo := userAppConfigProvider(c.App)
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
			repo := userAppConfigProvider(c.App)
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
