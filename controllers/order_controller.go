package controllers

import (
	"encoding/json"
	"fmt"
	web "github.com/igor-koniukhov/webLogger/v3"

	"github.com/igor-koniukhov/fastcat/internal/model"

	"github.com/igor-koniukhov/fastcat/internal/repository"
	"net/http"
)

type Order interface {
	Create() http.HandlerFunc
	Get() http.HandlerFunc
	GetAll() http.HandlerFunc
	Delete() http.HandlerFunc
	Update() http.HandlerFunc
}

type OrderController struct {
	repo repository.OrderRepositoryInterface
}

func NewOrderController(repo repository.OrderRepositoryInterface) *OrderController {
	return &OrderController{repo: repo}
}

func (ord OrderController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var o model.Order
		json.NewDecoder(r.Body).Decode(&ord)
		order, err := ord.repo.Create(&o)
		web.Log.Error(err, err)
		json.NewEncoder(w).Encode(&order)
	}
}

func (ord OrderController) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		id := param(r)
		order := ord.repo.Get(id)
		json.NewEncoder(w).Encode(&order)
	}
}

func (ord OrderController) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		order := ord.repo.GetAll()
		json.NewEncoder(w).Encode(&order)
	}
}

func (ord OrderController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := param(r)
		err := ord.repo.Delete(id)
		web.Log.Error(err, err)
		_, _ = fmt.Fprintf(w, fmt.Sprintf(" user with %d deleted", id))
	}
}

func (ord OrderController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var o model.Order
		_ = json.NewDecoder(r.Body).Decode(&ord)
		id := param(r)
		order := ord.repo.Update(id, &o)
		_ = json.NewEncoder(w).Encode(&order)
	}
}
