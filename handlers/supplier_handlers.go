package handlers

import (
	"encoding/json"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	web "github.com/igor-koniukhov/webLogger/v3"
	"net/http"
)

type Supplier interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type SupplierHandler struct {
	repo dbrepo.SupplierRepository
}

func NewSupplierHandler(repo dbrepo.SupplierRepository) *SupplierHandler {
	return &SupplierHandler{repo: repo}
}

func (s *SupplierHandler) Create(w http.ResponseWriter, r *http.Request) {
	var suppliers models.Suppliers
	json.NewDecoder(r.Body).Decode(&suppliers)
	user, err := s.repo.Create(&suppliers)
	web.Log.Error(err, err)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&user)
	web.Log.Error(err)
}

func (s *SupplierHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	id := param(r)
	user := s.repo.Get(id)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&user)
	web.Log.Error(err)
}

func (s *SupplierHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	suppliers := s.repo.GetAll()
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&suppliers)
	web.Log.Error(err)
}

func (s *SupplierHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := param(r)
	err := s.repo.Delete(id)
	web.Log.Error(err, err)
	w.WriteHeader(http.StatusAccepted)
}

func (s *SupplierHandler) Update(w http.ResponseWriter, r *http.Request) {
	var supplier models.Supplier
	json.NewDecoder(r.Body).Decode(&supplier)
	id := param(r)
	user := s.repo.Update(id, &supplier)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&user)
	web.Log.Error(err)
}
