package controllers

import (
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"net/http"
)

type CartControllerI interface {
	Create(method string) http.HandlerFunc
	Get(method string) http.HandlerFunc
	GetAll(method string) http.HandlerFunc
	DeleteC(method string) http.HandlerFunc
	Update(method string) http.HandlerFunc
}

var RepoCart *CartController

type CartController struct {
	App *config.AppConfig
}

func NewCartControllers(app *config.AppConfig) *CartController {
	return &CartController{App: app}
}
func NewControllersC(r *CartController)  {
	RepoCart = r

}
func cartAppConfigProvider(a *config.AppConfig) *repository.CartRepository {
	repo := repository.NewCartRepository(a)
	repository.NewRepoC(repo)
	return repo

}

func (c CartController) Create(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			cartAppConfigProvider(c.App)
		default:
			methodMessage(w, method)
		}
	}
}

func (c CartController) Get(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		switch r.Method {
		case method:
			cartAppConfigProvider(c.App)
		default:
			methodMessage(w, method)
		}
	}
}

func (c CartController) GetAll(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		switch r.Method {
		case method:
			cartAppConfigProvider(c.App)
		default:
			methodMessage(w, method)
		}
	}
}

func (c CartController) Delete(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			cartAppConfigProvider(c.App)
		default:
			methodMessage(w, method)
		}
	}
}

func (c CartController) Update(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			cartAppConfigProvider(c.App)
		default:
			methodMessage(w, method)
		}
	}
}
