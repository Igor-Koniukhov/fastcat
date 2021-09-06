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
	Create(method string) http.HandlerFunc
	Get(method string) http.HandlerFunc
	GetAllO(method string) http.HandlerFunc
	Delete(method string) http.HandlerFunc
	Update(method string) http.HandlerFunc
}

var RepoOrder *OrderController

type OrderController struct {
	App *config.AppConfig
}

func NewOrderControllers(app *config.AppConfig) *OrderController {
	return &OrderController{App: app}
}
func NewControllersO(r *OrderController) {
	RepoOrder = r
}
func orderAppConfigProvider(a *config.AppConfig) *repository.OrderRepository {
	repo := repository.NewOrderRepository(a)
	repository.NewRepoO(repo)
	return repo

}

func (ord OrderController) Create( method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var o model.Order
		switch r.Method {
		case method:

			json.NewDecoder(r.Body).Decode(&ord)
			orderAppConfigProvider(ord.App)
			order, err := repository.RepoO.CreateOrder(&o)
			checkError(err)
			json.NewEncoder(w).Encode(&order)

		default:
			methodMessage(w, method)
		}
	}
}


func (ord OrderController) Get(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		switch r.Method {
		case method:
			repo := orderAppConfigProvider(ord.App)
			param, _, _ := repo.Param(r)
			order := repository.RepoO.GetOrder(&param)
			json.NewEncoder(w).Encode(&order)
		default:
			methodMessage(w, method)
		}
	}
}

func (ord OrderController) GetAll(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		switch r.Method {
		case method:
			orderAppConfigProvider(ord.App)
			order := repository.RepoO.GetAllOrders()
			json.NewEncoder(w).Encode(&order)
		default:
			methodMessage(w, method)
		}
	}
}


func (ord OrderController) Delete(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			repo := orderAppConfigProvider(ord.App)
			_, _, id := repo.Param(r)
			err := repository.RepoO.DeleteOrder(id)
			checkError(err)
			_, _ = fmt.Fprintf(w, fmt.Sprintf(" user with %d deleted", id))

		default:
			methodMessage(w, method)
		}
	}
}

func (ord OrderController) Update(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			var o model.Order
			json.NewDecoder(r.Body).Decode(&ord)
			repo := orderAppConfigProvider(ord.App)
			_, _, id := repo.Param(r)
			order := repository.RepoO.UpdateOrder(id, &o)
			json.NewEncoder(w).Encode(&order)

		default:
			methodMessage(w, method)
		}
	}
}
