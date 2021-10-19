package handlers

import (
	"encoding/json"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/render"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	"github.com/igor-koniukhov/fastcat/services/router"
	web "github.com/igor-koniukhov/webLogger/v2"
	"log"
	"net/http"
)

type Order interface {
	ShowBlankOrder(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}
type OrderHandler struct {
	App  *config.AppConfig
	repo dbrepo.OrderRepository
}

func NewOrderHandler(app *config.AppConfig, repo dbrepo.OrderRepository) *OrderHandler {
	return &OrderHandler{App: app, repo: repo}
}
func (ord OrderHandler) ShowBlankOrder(w http.ResponseWriter, r *http.Request) {
	err := render.TemplateRender(w, r, "order.page.tmpl", &models.TemplateData{})
	if err != nil {
		web.Log.Fatal(err)
	}
	w.WriteHeader(http.StatusCreated)
}
func (ord OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	err := render.TemplateRender(w, r, "order.page.tmpl", &models.TemplateData{})
	if err != nil {
		web.Log.Fatal(err)
	}
	w.WriteHeader(http.StatusCreated)
}
func (ord OrderHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := router.GetKeyInt(r, ":id")
	order := ord.repo.Get(id)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&order)
	if err != nil {
		log.Println(err)
	}
}
func (ord OrderHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	order := ord.repo.GetAll()
	err := json.NewEncoder(w).Encode(&order)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}

func (ord OrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := router.GetKeyInt(r, ":id")
	err := ord.repo.Delete(id)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusAccepted)
}
func (ord OrderHandler) Update(w http.ResponseWriter, r *http.Request) {
	var o models.Order
	err := json.NewDecoder(r.Body).Decode(&ord)
	if err != nil {
		log.Println(err)
	}
	id := router.GetKeyInt(r, ":id")
	order := ord.repo.Update(id, &o)
	err = json.NewEncoder(w).Encode(&order)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}
