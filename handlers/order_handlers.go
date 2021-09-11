package handlers

import (
	"encoding/json"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	web "github.com/igor-koniukhov/webLogger/v3"

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
	repo dbrepo.OrderRepository
}

func NewOrderHandler(repo dbrepo.OrderRepository) *OrderController {
	return &OrderController{repo: repo}
}

func (ord OrderController) Create(w http.ResponseWriter, r *http.Request) {
	var o models.Order
	json.NewDecoder(r.Body).Decode(&o)
	order, err := ord.repo.Create(&o)
	web.Log.Error(err, err)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&order)
	web.Log.Error(err)
}

func (ord OrderController) Get(w http.ResponseWriter, r *http.Request) {
	id := param(r)
	order := ord.repo.Get(id)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&order)
	web.Log.Error(err)
}

func (ord OrderController) GetAll(w http.ResponseWriter, r *http.Request) {
	order := ord.repo.GetAll()
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&order)
	web.Log.Error(err)
}

func (ord OrderController) Delete(w http.ResponseWriter, r *http.Request) {
	id := param(r)
	err := ord.repo.Delete(id)
	w.WriteHeader(http.StatusAccepted)
	web.Log.Error(err, err)
}

func (ord OrderController) Update(w http.ResponseWriter, r *http.Request) {
	var o models.Order
	_ = json.NewDecoder(r.Body).Decode(&ord)
	id := param(r)
	order := ord.repo.Update(id, &o)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&order)
	web.Log.Error(err)
}
