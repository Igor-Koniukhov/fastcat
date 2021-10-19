package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/render"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	"github.com/igor-koniukhov/fastcat/services/router"
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
	GetAllByUserID(w http.ResponseWriter, r *http.Request)
}

type CartHandler struct {
	App  *config.AppConfig
	repo dbrepo.CartRepository
}

func NewCartHandler(app *config.AppConfig, repo dbrepo.CartRepository) *CartHandler {
	return &CartHandler{App: app, repo: repo}
}
func (c CartHandler) Create(w http.ResponseWriter, r *http.Request) {
	str := r.Context().Value("user_id")
	userId := str.(int)

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
		Tel:   r.Form.Get("phone"),
	}
	userString, err := json.Marshal(user)
	if err != nil {
		web.Log.Fatal(err)
	}
	or := &models.CartResponse{
		UserID:          userId,
		User:            userString,
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
	id := router.GetKeyInt(r, ":id")
	cart := c.repo.Get(id)
	err := render.TemplateRender(w, r, "show_user_cart.page.tmpl", &models.TemplateData{Cart: cart})
	if err != nil {
		web.Log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
}
func (c CartHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	cart := c.repo.GetAll()
	err := json.NewEncoder(w).Encode(&cart)
	if err != nil {
		web.Log.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (c CartHandler) GetAllByUserID(w http.ResponseWriter, r *http.Request){
	userId := r.Context().Value("user_id")
	id := userId.(int)
	carts, err := c.repo.GetAllByUserID(id)
	if err !=nil {
		web.Log.Error(err)
	}
	err = render.TemplateRender(w, r, "user_cabinet.page.tmpl",
		&models.TemplateData{
			UserCabinetInfo: carts,
		})
	if err != nil {
		web.Log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
}

func (c CartHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := router.GetKeyInt(r, ":id")
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
	id := router.GetKeyInt(r, ":id")
	crt := c.repo.Update(id, &cart)
	err = json.NewEncoder(w).Encode(&crt)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}
