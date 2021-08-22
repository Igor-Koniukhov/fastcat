package controllers

import (
	"github.com/igor-koniukhov/fastcat/internal/config"
	"net/http"
)

type CartControllerI interface {
	CreateCart(method string) http.HandlerFunc
	GetCart(method string) http.HandlerFunc
	GetAllCarts(method string) http.HandlerFunc
	DeleteCarts(method string) http.HandlerFunc
	UpdateCart(method string) http.HandlerFunc
}

var RepoCart *CartControllers

type CartControllers struct {
	App *config.AppConfig
}

func NewCartControllers(app *config.AppConfig) *CartControllers {
	return &CartControllers{App: app}
}
func NewControllersC(r *CartControllers)  {
	RepoCart = r

}

func (c CartControllers) CreateCart(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:

		default:
			methodMassage(w, method)
		}
	}
}

func (c CartControllers) GetCart(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:

		default:
			methodMassage(w, method)
		}
	}
}

func (c CartControllers) GetAllCarts(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:

		default:
			methodMassage(w, method)
		}
	}
}

func (c CartControllers) DeleteCart(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:

		default:
			methodMassage(w, method)
		}
	}
}

func (c CartControllers) UpdateCart(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:

		default:
			methodMassage(w, method)
		}
	}
}
