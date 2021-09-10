package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	web "github.com/igor-koniukhov/webLogger/v3"
	"net/http"
)

type Supplier interface {
	Create() http.HandlerFunc
	Get() http.HandlerFunc
	GetAll() http.HandlerFunc
	Delete() http.HandlerFunc
	Update() http.HandlerFunc
}

type SupplierController struct {
	repo dbrepo.SupplierRepository
}

func NewSupplierController(repo dbrepo.SupplierRepository) *SupplierController {
	return &SupplierController{repo: repo}
}

func (s SupplierController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var suppliers model.Suppliers
		json.NewDecoder(r.Body).Decode(&suppliers)
		user, err := s.repo.Create(&suppliers)
		web.Log.Error(err, err)
		json.NewEncoder(w).Encode(&user)
	}
}

func (s SupplierController) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		id := param(r)
		user := s.repo.Get(id)
		json.NewEncoder(w).Encode(&user)
	}
}

func (s SupplierController) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		suppliers := s.repo.GetAll()
		json.NewEncoder(w).Encode(&suppliers)
	}
}

func (s SupplierController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := param(r)
		err := s.repo.Delete(id)
		web.Log.Error(err, err)
		_, _ = fmt.Fprintf(w, fmt.Sprintf(" supplier with %d deleted", id))
	}
}

func (s SupplierController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var supplier model.Supplier
		json.NewDecoder(r.Body).Decode(&supplier)
		id := param(r)
		user := s.repo.Update(id, &supplier)
		json.NewEncoder(w).Encode(&user)
	}
}
