package handlers

import (
	"encoding/json"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	web "github.com/igor-koniukhov/webLogger/v3"
	"net/http"
	"strconv"
	"strings"
)

type Product interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type ProductController struct {
	repo dbrepo.ProductRepository
}

func NewProductHandler(repo dbrepo.ProductRepository) *ProductController {
	return &ProductController{repo: repo}
}

func (p *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

func param(r *http.Request) (id int) {
	fields := strings.Split(r.URL.String(), "/")
	str := fields[len(fields)-1]
	id, err := strconv.Atoi(str)
	web.Log.Error(err, err)
	return
}

func (p *ProductController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	id := param(r)
	item := p.repo.Get(id)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&item)
	web.Log.Error(err)
}

func (p *ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	items := p.repo.GetAll()
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&items)
	web.Log.Error(err)
}

func (p *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	id := param(r)
	err := p.repo.Delete(id)
	web.Log.Error(err, err)
	w.WriteHeader(http.StatusAccepted)
}

func (p *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	var item *models.Item
	json.NewDecoder(r.Body).Decode(&item)
	id := param(r)
	item = p.repo.Update(id, item)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&item)
	web.Log.Error(err)

}
