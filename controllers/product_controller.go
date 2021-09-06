package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	web "github.com/igor-koniukhov/webLogger/v3"
	"net/http"
)

type ProductControllerI interface {
	Create(method string) http.HandlerFunc
	Get(method string) http.HandlerFunc
	GetAllP(method string) http.HandlerFunc
	Delete(method string) http.HandlerFunc
	Update(method string) http.HandlerFunc
}

var RepoProducts *ProductController

type ProductController struct {
	App *config.AppConfig
}

func NewProductControllers(app *config.AppConfig) *ProductController {
	return &ProductController{App: app}
}
func NewControllersP(r *ProductController) {
	RepoProducts = r
}
func productAppConfigProvider(a *config.AppConfig) *repository.ProductRepository {
	repo := repository.NewProductRepository(a)
	repository.NewRepoP(repo)
	return repo

}

func (p *ProductController) Create(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			productAppConfigProvider(p.App)

		default:
			methodMessage(w, method)
		}
	}
}

func (p *ProductController) Get(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		switch r.Method {
		case method:
			repo := productAppConfigProvider(p.App)
			_, _, id := repo.Param(r)
			item := repository.RepoP.Get(id)

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
			 productAppConfigProvider(p.App)
			items := repository.RepoP.GetAll()
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
			repo := productAppConfigProvider(p.App)
			_, _, id := repo.Param(r)
			err := repository.RepoP.Delete(id)
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
			productAppConfigProvider(p.App)
			var item *model.Item
			json.NewDecoder(r.Body).Decode(&item)
			repo := supplierAppConfigProvider(p.App)
			_, _, id := repo.Param(r)
			item = repository.RepoP.Update(id, item)
			json.NewEncoder(w).Encode(&item)
		default:
			methodMessage(w, method)
		}
	}
}
