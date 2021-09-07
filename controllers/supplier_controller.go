package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	web "github.com/igor-koniukhov/webLogger/v3"
	"net/http"
)

type Supplier interface {
	Create(method string) http.HandlerFunc
	Get(method string) http.HandlerFunc
	GetAll(method string) http.HandlerFunc
	Delete(method string) http.HandlerFunc
	Update(method string) http.HandlerFunc
}

type SupplierController struct{
	repo repository.SupplierRepositoryInterface
}

func NewSupplierController(repo repository.SupplierRepositoryInterface) *SupplierController {
	return &SupplierController{repo: repo}
}

func (s SupplierController) Create(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			var suppliers model.Suppliers
			json.NewDecoder(r.Body).Decode(&suppliers)
			user, err := s.repo.Create(&suppliers)
			web.Log.Error(err, err)
			json.NewEncoder(w).Encode(&user)
		default:
			methodMessage(w, method)
		}
	}
}

func (s SupplierController) Get(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		switch r.Method {
		case method:
			param, nameParam, _ := s.repo.Param(r)
			user := s.repo.Get(&nameParam, &param)
			json.NewEncoder(w).Encode(&user)
		default:
			methodMessage(w, method)
		}
	}
}

func (s SupplierController) GetAll(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		switch r.Method {
		case method:
			suppliers := s.repo.GetAll()
			json.NewEncoder(w).Encode(&suppliers)

			methodMessage(w, method)
		}
	}
}

func (s SupplierController) Delete(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			_, _, id := s.repo.Param(r)
			err := s.repo.Delete(id)
			web.Log.Error(err, err)
			_, _ = fmt.Fprintf(w, fmt.Sprintf(" supplier with %d deleted", id))
		default:
			methodMessage(w, method)
		}
	}
}

func (s SupplierController) Update(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			var supplier model.Supplier
			json.NewDecoder(r.Body).Decode(&supplier)
			_, _, id := s.repo.Param(r)
			user := s.repo.Update(id, &supplier)
			json.NewEncoder(w).Encode(&user)
		default:
			methodMessage(w, method)
		}
	}
}
