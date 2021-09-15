package handlers

import (
	"encoding/json"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	"log"
	"net/http"
)

type Cart interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type CartHandler struct {
	App *config.AppConfig
	repo dbrepo.CartRepository
}

func NewCartHandler(app *config.AppConfig, repo dbrepo.CartRepository) *CartHandler {
	return &CartHandler{App: app, repo: repo}
}

func (c CartHandler) Create(w http.ResponseWriter, r *http.Request){
	var cart models.Cart
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		log.Println(err)
	}
	crt, err := c.repo.Create(&cart)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&crt)
	if err != nil {
		log.Println(err)
	}
}

func (c CartHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	id := param(r)
	cart := c.repo.Get(id)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&cart)
	if err != nil {
		log.Println(err)
	}
}

func (c CartHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	cart := c.repo.GetAll()
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&cart)
	if err != nil {
		log.Println(err)
	}
}

func (c CartHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := param(r)
	err := c.repo.Delete(id)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusAccepted)
}

func (c CartHandler) Update(w http.ResponseWriter, r *http.Request) {
	var cart models.Cart
	err:= json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		log.Println(err)
	}
	id := param(r)
	crt := c.repo.Update(id, &cart)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&crt)
	if err != nil {
		log.Println(err)
	}
}
