package handlers

import (
	"encoding/json"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/render"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	"github.com/igor-koniukhov/fastcat/services/router"
	web "github.com/igor-koniukhov/webLogger/v2"
	"log"
	"net/http"
	"strconv"
)

type Product interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetAllBySupplierID(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetJson(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type ProductHandler struct {
	App  *config.AppConfig
	repo dbrepo.ProductRepository
}

func NewProductHandler(app *config.AppConfig, repo dbrepo.ProductRepository) *ProductHandler {
	return &ProductHandler{App: app, repo: repo}
}
func (p *ProductHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := router.GetKeyInt(r, ":id")
	item := p.repo.Get(id)
	err := json.NewEncoder(w).Encode(&item)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}
func (p *ProductHandler) GetAllBySupplierID(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		web.Log.Error(err)
	}
	id, err := strconv.Atoi(r.Form.Get("supplier_id"))
	if err != nil {
		web.Log.Error(err)
	}
	supplier := repository.Repo.SupplierRepository.Get(id)
	products := p.repo.GetAllBySupplierID(id)
	err = render.TemplateRender(w, r, "products.page.tmpl",
		&models.TemplateData{
			Products:     products,
			Supplier:     supplier,
		})
	if err != nil {
		web.Log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
}
func (p *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	items := p.repo.GetAll()
	var supp []models.Supplier
	for i := 0; i < len(items)-1; i++ {
		s := repository.Repo.SupplierRepository.Get(items[i].SuppliersID)
		supp = append(supp, *s)
	}
	err := render.TemplateRender(w, r, "products.page.tmpl",
		&models.TemplateData{
			Products:  items,
			Suppliers: supp,
		})
	if err != nil {
		web.Log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
}
func (p *ProductHandler) GetJson(w http.ResponseWriter, r *http.Request) {
	items := p.repo.GetAll()
	err := json.NewEncoder(w).Encode(&items)
	if err != nil {
		web.Log.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (p *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := router.GetKeyInt(r, ":id")
	err := p.repo.Delete(id)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusAccepted)
}
func (p *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	var item *models.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(err)
	}
	id := router.GetKeyInt(r, ":id")
	item = p.repo.Update(id, item)
	err = json.NewEncoder(w).Encode(&item)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}
