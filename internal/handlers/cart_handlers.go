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
	App  *config.AppConfig
	repo dbrepo.CartRepository
}

func NewCartHandler(app *config.AppConfig, repo dbrepo.CartRepository) *CartHandler {
	return &CartHandler{App: app, repo: repo}
}

func (c CartHandler) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		web.Log.Fatal(err)
	}
	cartCookie, err := r.Cookie("Product")
	if err != nil {
		log.Fatal(err)
		return
	}
	cartBody, err := url.QueryUnescape(cartCookie.Value)
	sb := []byte(cartBody)
	var cb []models.CartBody
	json.Unmarshal(sb, &cb)
	for _, v := range cb {
		fmt.Println(v.Title, v.ProductID, v.SupplierID, v.Price)
	}
	if err != nil {
		log.Fatal(err)
		return
	}
	p := r.Form.Get("prodInfo")
	user := &models.User{
		Name:  r.Form.Get("name"),
		Email: r.Form.Get("email"),
		Phone: r.Form.Get("phone"),
	}
	userB, err := json.Marshal(user)
	if err != nil {
		web.Log.Fatal(err)
	}
	or := &models.CartResponse{
		User:            userB,
		AddressDelivery: r.Form.Get("address"),
		CartBody:        []byte(p),
		Amount:          r.Form.Get("amount"),
	}
	_, id, err := c.repo.Create(or)
	if err != nil {
		web.Log.Fatal(err)
	}
	url := fmt.Sprintf("/cart/%d", id)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (c CartHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := param(r)
	cart := c.repo.Get(id)
	w.WriteHeader(http.StatusOK)
	err := render.TemplateRender(w, r, "show_user_cart.page.tmpl", &models.TemplateData{Cart: cart})
	if err != nil {
		web.Log.Fatal(err)
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
	err := json.NewDecoder(r.Body).Decode(&cart)
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
