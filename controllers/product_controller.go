package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	web "github.com/igor-koniukhov/webLogger/v3"
	"net/http"
	"strconv"
	"strings"
)

type Product interface {
	Create() http.HandlerFunc
	Get() http.HandlerFunc
	GetAll() http.HandlerFunc
	Delete() http.HandlerFunc
	Update() http.HandlerFunc
}

type ProductController struct {
	repo dbrepo.ProductRepository
}

func NewProductController(repo dbrepo.ProductRepository) *ProductController {
	return &ProductController{repo: repo}
}

func (p *ProductController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func  param(r *http.Request) (id int) {
	fields := strings.Split(r.URL.String(), "/")
	str := fields[len(fields)-1]
		id, err := strconv.Atoi(str)
		web.Log.Error(err, err)
	return
}


func (p *ProductController) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		id := param(r)
		item := p.repo.Get(id)
		json.NewEncoder(w).Encode(&item)
	}
}

func (p *ProductController) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		items := p.repo.GetAll()
		json.NewEncoder(w).Encode(&items)
	}
}

func (p *ProductController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := param(r)
		err := p.repo.Delete(id)
		web.Log.Error(err, err)
		_, _ = fmt.Fprintf(w, fmt.Sprintf(" product with %d deleted", id))
	}
}

func (p *ProductController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item *model.Item
		json.NewDecoder(r.Body).Decode(&item)
		id := param(r)
		item = p.repo.Update(id, item)
		json.NewEncoder(w).Encode(&item)
	}
}
