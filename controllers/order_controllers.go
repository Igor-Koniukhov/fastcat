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

func (oc OrderControllers) CreateOrder( method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ord model.Order
		switch r.Method {
		case method:
			json.NewDecoder(r.Body).Decode(&ord)
			orderRepo := repository.OrderRepository{}
			order, err := orderRepo.CreateOrder(&ord, oc.App.DB)
			checkError(err)
			json.NewEncoder(w).Encode(&order)
		default:
			methodMassage(w, method)
		}
	}
}

func (oc OrderControllers) GetOrder(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(oc.App.Session)

		switch r.Method {
		case method:
			orderRepo := repository.OrderRepository{}
			param, _, _ := orderRepo.Param(r)
			order := orderRepo.GetOrder(&param, oc.App.DB)
			json.NewEncoder(w).Encode(&order)
		default:
			methodMassage(w, method)
		}
	}
}

func (oc OrderControllers) GetAllOrders(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			orderRepo := repository.OrderRepository{}
			order := orderRepo.GetAllOrders(oc.App.DB)
			json.NewEncoder(w).Encode(&order)
		default:
			methodMassage(w, method)
		}
	}
}

func (oc OrderControllers) DeleteOrder(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			orderRepo := repository.OrderRepository{}
			_, _, id := orderRepo.Param(r)
			err := orderRepo.DeleteOrder(id, oc.App.DB)
			checkError(err)
			_, _ = fmt.Fprintf(w, fmt.Sprintf(" user with %d deleted", id))
		default:
			methodMassage(w, method)
		}
	}
}

func (oc OrderControllers) UpdateOrder(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			var ord model.Order
			json.NewDecoder(r.Body).Decode(&ord)
			orderRepo := repository.OrderRepository{}
			_, _, id := orderRepo.Param(r)
			order := orderRepo.UpdateOrder(id, &ord, oc.App.DB)
			json.NewEncoder(w).Encode(&order)

		default:
			methodMassage(w, method)
		}
	}
}
