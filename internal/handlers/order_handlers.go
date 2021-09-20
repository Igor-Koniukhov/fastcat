package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/render"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	web "github.com/igor-koniukhov/webLogger/v2"
	"log"
	"net/url"

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
	w.WriteHeader(http.StatusCreated)
	err := render.TemplateRender(w, r, "order.page.tmpl", models.TemplateData{})
	if err != nil {
		web.Log.Fatal(err)
	}
}

func (ord OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		web.Log.Fatal(err)
	}
	var prods []models.Prod
	orderCookie, err := r.Cookie("Product")
	if err != nil {
		log.Fatal(err)
		return
	}
	decodedValue, err := url.QueryUnescape(orderCookie.Value)
	err = json.Unmarshal([]byte(decodedValue), &prods)
	for _, v := range prods {
		fmt.Println(v.Title, v.Price)
	}
	if err != nil {
		log.Fatal(err)
		return
	}
	user := &models.User{
		Name:  r.Form.Get("name"),
		Email: r.Form.Get("email"),
		Phone: r.Form.Get("phone"),
	}
	web.Log.Info(user, prods)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (ord OrderHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := param(r)
	order := ord.repo.Get(id)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&order)
	if err != nil {
		log.Println(err)
	}
}

func (ord OrderHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	order := ord.repo.GetAll()
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&order)
	if err != nil {
		log.Println(err)
	}
}

func (ord OrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := param(r)
	err := ord.repo.Delete(id)
	w.WriteHeader(http.StatusAccepted)
	if err != nil {
		log.Println(err)
	}
}

func (ord OrderHandler) Update(w http.ResponseWriter, r *http.Request) {
	var o models.Order
	err := json.NewDecoder(r.Body).Decode(&ord)
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
