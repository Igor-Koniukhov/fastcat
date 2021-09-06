package controllers

import (
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"net/http"
)

type CartControllerI interface {
	CreateCart(method string) http.HandlerFunc
	GetCart(method string) http.HandlerFunc
	GetAllCarts(method string) http.HandlerFunc
	DeleteCarts(method string) http.HandlerFunc
	UpdateCart(method string) http.HandlerFunc
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

func (c CartController) CreateCart(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			cartAppConfigProvider(c.App)
		default:
			methodMessage(w, method)
		}
	}
}

func (c CartController) GetCart(method string) http.HandlerFunc {
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

func (c CartController) GetAllCarts(method string) http.HandlerFunc {
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

func (c CartController) DeleteCart(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			cartAppConfigProvider(c.App)
		default:
			methodMessage(w, method)
		}
	}
}

func (c CartController) UpdateCart(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			cartAppConfigProvider(c.App)
		default:
			methodMessage(w, method)
		}
	}
}
