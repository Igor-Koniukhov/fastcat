package handlers

import (
	"encoding/json"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	"log"

	"github.com/igor-koniukhov/fastcat/internal/models"

	"net/http"
)

type Order interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type OrderController struct {
	App *config.AppConfig
	repo dbrepo.OrderRepository
}

func NewOrderHandler(app *config.AppConfig, repo dbrepo.OrderRepository) *OrderController {
	return &OrderController{App: app, repo: repo}
}

func (ord OrderController) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)



}

func (ord OrderController) Get(w http.ResponseWriter, r *http.Request) {
	id := param(r)
	order := ord.repo.Get(id)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&order)
	if err != nil {
		log.Println(err)
	}
}

func (ord OrderController) GetAll(w http.ResponseWriter, r *http.Request) {
	order := ord.repo.GetAll()
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&order)
	if err != nil {
		log.Println(err)
	}
}

func (ord OrderController) Delete(w http.ResponseWriter, r *http.Request) {
	id := param(r)
	err := ord.repo.Delete(id)
	w.WriteHeader(http.StatusAccepted)
	if err != nil {
		log.Println(err)
	}
}

func (ord OrderController) Update(w http.ResponseWriter, r *http.Request) {
	var o models.Order
	err:= json.NewDecoder(r.Body).Decode(&ord)
	if err != nil {
		log.Println(err)
	}
	id := param(r)
	order := ord.repo.Update(id, &o)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&order)
	if err != nil {
		log.Println(err)
	}
}
