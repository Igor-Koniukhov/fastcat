package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"net/http"
)

type OrderControllerI interface {
	CreateOrder(method string) http.HandlerFunc
	GetOrder(method string) http.HandlerFunc
	GetAllOrders(method string) http.HandlerFunc
	DeleteOrder(method string) http.HandlerFunc
	UpdateOrder(method string) http.HandlerFunc
}

var RepoOrder *OrderControllers

type OrderControllers struct {
	App *config.AppConfig
}

func NewOrderControllers(app *config.AppConfig) *OrderControllers {
	return &OrderControllers{App: app}
}
func NewControllersO(r *OrderControllers) {
	RepoOrder = r
}

func (o OrderControllers) CreateOrder(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ord model.Order
		switch r.Method {
		case method:
			json.NewDecoder(r.Body).Decode(&ord)
			orderRepo := repository.OrderRepository{}
			order, err := orderRepo.CreateOrder(&ord, o.App.DB)
			checkError(err)
			json.NewEncoder(w).Encode(&order)
		default:
			methodMassage(w, method)
		}
	}
}

func (o OrderControllers) GetOrder(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			orderRepo := repository.OrderRepository{}
			param, _, _ := orderRepo.Param(r)
			order := orderRepo.GetOrder(&param, o.App.DB)
			json.NewEncoder(w).Encode(&order)
		default:
			methodMassage(w, method)
		}
	}
}

func (o OrderControllers) GetAllOrders(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			orderRepo := repository.OrderRepository{}
			order := orderRepo.GetAllOrders(o.App.DB)
			json.NewEncoder(w).Encode(&order)
		default:
			methodMassage(w, method)
		}
	}
}

func (o OrderControllers) DeleteOrder(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			orderRepo := repository.OrderRepository{}
			_, _, id := orderRepo.Param(r)
			err := orderRepo.DeleteOrder(id, o.App.DB)
			checkError(err)
			_, _ = fmt.Fprintf(w, fmt.Sprintf(" user with %d deleted", id))
		default:
			methodMassage(w, method)
		}
	}
}

func (o OrderControllers) UpdateOrder(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			var ord model.Order
			json.NewDecoder(r.Body).Decode(&ord)
			orderRepo := repository.OrderRepository{}
			_, _, id := orderRepo.Param(r)
			order := orderRepo.UpdateOrder(id, &ord, o.App.DB)
			json.NewEncoder(w).Encode(&order)

		default:
			methodMassage(w, method)
		}
	}
}
