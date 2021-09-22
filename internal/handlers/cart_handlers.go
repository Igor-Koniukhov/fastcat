package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	web "github.com/igor-koniukhov/webLogger/v2"
	"log"
	"net/http"
	"net/url"
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
	err := r.ParseForm()
	if err != nil {
		web.Log.Fatal(err)
	}
	orderCookie, err := r.Cookie("Product")
	if err != nil {
		log.Fatal(err)
		return
	}
	orderBody, err := url.QueryUnescape(orderCookie.Value)
	if err != nil {
		log.Fatal(err)
		return
	}
	p:= r.Form.Get("prodInfo")
	fmt.Println(p, "this is prodinfo")
	user := &models.User{
		Name:  r.Form.Get("name"),
		Email: r.Form.Get("email"),
		Phone: r.Form.Get("phone"),
	}
	or := &models.CartResponse{
		User:            *user,
		AddressDelivery: r.Form.Get("address"),
		OrderBody:     orderBody,
		Amount:          r.Form.Get("amount"),
	}
	web.Log.Info( or)
	//http.Redirect(w, r, "/products", http.StatusSeeOther)
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
	var cart models.CartResponse
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
