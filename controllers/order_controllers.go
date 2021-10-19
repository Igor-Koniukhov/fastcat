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
func orderAppConfigProvider(a *config.AppConfig) *repository.OrderRepository {
	repo := repository.NewOrderRepository(a)
	repository.NewRepoO(repo)
	return repo

}

func (o OrderControllers) CreateOrder( method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ord model.Order
		switch r.Method {
		case method:

			json.NewDecoder(r.Body).Decode(&ord)
			orderAppConfigProvider(o.App)
			order, err := repository.RepoO.CreateOrder(&ord)
			checkError(err)
			json.NewEncoder(w).Encode(&order)

		default:
			methodMessage(w, method)
		}
	}
}


func (o OrderControllers) GetOrder(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		switch r.Method {
		case method:
			repo := orderAppConfigProvider(o.App)
			param, _, _ := repo.Param(r)
			order := repository.RepoO.GetOrder(&param)
			json.NewEncoder(w).Encode(&order)
		default:
			methodMessage(w, method)
		}
	}
}

func (o OrderControllers) GetAllOrders(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		switch r.Method {
		case method:
			orderAppConfigProvider(o.App)
			order := repository.RepoO.GetAllOrders()
			json.NewEncoder(w).Encode(&order)
		default:
			methodMessage(w, method)
		}
	}
}


func (o OrderControllers) DeleteOrder(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			repo := orderAppConfigProvider(o.App)
			_, _, id := repo.Param(r)
			err := repository.RepoO.DeleteOrder(id)
			checkError(err)
			_, _ = fmt.Fprintf(w, fmt.Sprintf(" user with %d deleted", id))

		default:
			methodMessage(w, method)
		}
	}
}

func (oc OrderControllers) UpdateOrder(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			var ord model.Order
			json.NewDecoder(r.Body).Decode(&ord)
			repo := orderAppConfigProvider(oc.App)
			_, _, id := repo.Param(r)
			order := repository.RepoO.UpdateOrder(id, &ord)
			json.NewEncoder(w).Encode(&order)

		default:
			methodMessage(w, method)
		}
	}
}
