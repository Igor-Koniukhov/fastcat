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
	Create(method string) http.HandlerFunc
	Get(method string) http.HandlerFunc
	GetAll(method string) http.HandlerFunc
	Delete(method string) http.HandlerFunc
	Update(method string) http.HandlerFunc
}

type OrderController struct {
	repo repository.OrderRepositoryInterface
}

func (ord OrderController) Create( method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var o model.Order
		switch r.Method {
		case method:
			json.NewDecoder(r.Body).Decode(&ord)
			order, err := ord.repo.Create(&o)
			web.Log.Error(err, err)
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
			param, _, _ := ord.repo.Param(r)
			order := ord.repo.Get(&param)
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
			order := ord.repo.GetAll()
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
			_, _, id := ord.repo.Param(r)
			err := ord.repo.Delete(id)
			web.Log.Error(err, err)
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
			_, _, id := ord.repo.Param(r)
			order := ord.repo.Update(id, &o)
			json.NewEncoder(w).Encode(&order)

		default:
			methodMessage(w, method)
		}
	}
}
