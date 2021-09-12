package handlers

import (
	"encoding/json"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	web "github.com/igor-koniukhov/webLogger/v3"
	"net/http"
)

type Cart interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type CartController struct {
	repo dbrepo.CartRepository
}

func NewCartHandler(repo dbrepo.CartRepository) *CartController {
	return &CartController{repo: repo}
}

func (c CartController) Create(w http.ResponseWriter, r *http.Request){
	var cart models.Cart
	json.NewDecoder(r.Body).Decode(&cart)
	crt, err := c.repo.Create(&cart)
	web.Log.Error(err, err)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&crt)
	web.Log.Error(err)
}

func (c CartController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	id := param(r)
	cart := c.repo.Get(id)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&cart)
	web.Log.Error(err)
}

func (c CartController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	cart := c.repo.GetAll()
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&cart)
	web.Log.Error(err)
}

func (c CartController) Delete(w http.ResponseWriter, r *http.Request) {
	id := param(r)
	err := c.repo.Delete(id)
	web.Log.Error(err, err)
	w.WriteHeader(http.StatusAccepted)
}

func (c CartController) Update(w http.ResponseWriter, r *http.Request) {
	var cart models.Cart
	_ = json.NewDecoder(r.Body).Decode(&cart)
	id := param(r)
	crt := c.repo.Update(id, &cart)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&crt)
	web.Log.Error(err)
}
