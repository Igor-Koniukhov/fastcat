package controllers

import (
	"github.com/igor-koniukhov/fastcat/internal/config"
	"net/http"
)

type OrderControllerI interface {
	CreateOrder( method string) http.HandlerFunc
	GetOrder( method string) http.HandlerFunc
	GetAllOrders( method string) http.HandlerFunc
	DeleteOrder( method string) http.HandlerFunc
	UpdateOrder( method string) http.HandlerFunc
}
var RepoOrder *OrderControllers
type OrderControllers struct {
	App *config.AppConfig
}

func NewOrderControllers(app *config.AppConfig) *OrderControllers {
	return &OrderControllers{App: app}
}
func NewControllersO(r *OrderControllers)  {
	RepoOrder = r
}

func (o OrderControllers) CreateOrder( method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

func (o OrderControllers) GetOrder( method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

func (o OrderControllers) GetAllOrders( method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

func (o OrderControllers) DeleteOrder( method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

func (o OrderControllers) UpdateOrder( method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}
