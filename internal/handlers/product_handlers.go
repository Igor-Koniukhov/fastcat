package handlers

import (
	"encoding/json"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/render"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Product interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetByTime(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
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

func (p *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

func param(r *http.Request) (id int) {
	fields := strings.Split(r.URL.String(), "/")
	str := fields[len(fields)-1]
	id, err := strconv.Atoi(str)
	if err != nil {
		log.Println(err)
	}
	return
}

func (p *ProductHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	id := param(r)
	item := p.repo.Get(id)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&item)
	if err != nil {
		log.Println(err)
	}
}
func (p *ProductHandler) GetByTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	id := param(r)
	item := p.repo.Get(id)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&item)
	if err != nil {
		log.Println(err)
	}
}

func (p *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	setPCook := &http.Cookie{
		Name:  "Product",
		Value: "",
	}
	http.SetCookie(w, setPCook)
	items := p.repo.GetAll()
	var supp []models.Supplier
	for i := 0; i < len(items)-1; i++ {
		s := repository.Repo.SupplierRepository.Get(items[i].SuppliersID)
		supp = append(supp, *s)
	}

	w.WriteHeader(http.StatusOK)
	render.TemplateRender(w, r, "product.page.tmpl", models.TemplateData{Products: items, Suppliers: supp})
}

func (p *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := param(r)
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
	id := param(r)
	item = p.repo.Update(id, item)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&item)
	if err != nil {
		log.Println(err)
	}

}
