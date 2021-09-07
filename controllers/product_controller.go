package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	web "github.com/igor-koniukhov/webLogger/v3"
	"net/http"
)

type Product interface {
	Create(method string) http.HandlerFunc
	Get(method string) http.HandlerFunc
	GetAllP(method string) http.HandlerFunc
	Delete(method string) http.HandlerFunc
	Update(method string) http.HandlerFunc
}


type ProductController struct {
	repo repository.ProductRepositoryInterface
}

func NewProductController(repo repository.ProductRepositoryInterface) *ProductController {
	return &ProductController{repo: repo}
}

func (p *ProductController) Create(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (p *ProductController) Get(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		switch r.Method {
		case method:
			_, _, id := p.repo.Param(r)
			item := p.repo.Get(id)
			json.NewEncoder(w).Encode(&item)
		default:
			methodMessage(w, method)
		}
	}
}

func (p *ProductController) GetAll(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		switch r.Method {
		case method:
			items := p.repo.GetAll()
			json.NewEncoder(w).Encode(&items)
		default:
			methodMessage(w, method)
		}
	}
}

func (p *ProductController) Delete(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
				_, _, id := p.repo.Param(r)
			err := p.repo.Delete(id)
			web.Log.Error(err, err)
			_, _ = fmt.Fprintf(w, fmt.Sprintf(" product with %d deleted", id))
		default:
			methodMessage(w, method)
		}
	}
}

func (p *ProductController) Update(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			var item *model.Item
			json.NewDecoder(r.Body).Decode(&item)
			_, _, id := p.repo.Param(r)
			item = p.repo.Update(id, item)
			json.NewEncoder(w).Encode(&item)
		default:
			methodMessage(w, method)
		}
	}
}
