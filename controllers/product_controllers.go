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
	CreateProduct(method string) http.HandlerFunc
	GetProduct(method string) http.HandlerFunc
	GetAllProducts(method string) http.HandlerFunc
	DeleteProduct(method string) http.HandlerFunc
	UpdateProduct(method string) http.HandlerFunc
}

var RepoProducts *ProductControllers

type ProductControllers struct {
	App *config.AppConfig
}

func NewProductControllers(app *config.AppConfig) *ProductControllers {
	return &ProductControllers{App: app}
}
func NewControllersP(r *ProductControllers) {
	RepoProducts = r
}
func productAppConfigProvider(a *config.AppConfig) *repository.ProductRepository {
	repo := repository.NewProductRepository(a)
	repository.NewRepoP(repo)
	return repo

}

func (p *ProductControllers) CreateProduct(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			productAppConfigProvider(p.App)

		default:
			methodMessage(w, method)
		}
	}
}

func (p *ProductControllers) GetProduct(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		switch r.Method {
		case method:
			repo := productAppConfigProvider(p.App)
			_, _, id := repo.Param(r)
			item := repository.RepoP.GetProduct(id)

			json.NewEncoder(w).Encode(&item)
		default:
			methodMessage(w, method)
		}
	}
}

func (p *ProductControllers) GetAllProducts(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		switch r.Method {
		case method:
			 productAppConfigProvider(p.App)
			items := repository.RepoP.GetAllProducts()
			json.NewEncoder(w).Encode(&items)
		default:
			methodMessage(w, method)
		}
	}
}

func (p *ProductControllers) DeleteProduct(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			repo := productAppConfigProvider(p.App)
			_, _, id := repo.Param(r)
			err := repository.RepoP.DeleteProduct(id)
			web.Log.Error(err, err)
			_, _ = fmt.Fprintf(w, fmt.Sprintf(" product with %d deleted", id))
		default:
			methodMessage(w, method)
		}
	}
}

func (p *ProductControllers) UpdateProduct(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			productAppConfigProvider(p.App)
			var item *model.Item
			json.NewDecoder(r.Body).Decode(&item)
			repo := supplierAppConfigProvider(p.App)
			_, _, id := repo.Param(r)
			item = repository.RepoP.UpdateProduct(id, item)
			json.NewEncoder(w).Encode(&item)
		default:
			methodMessage(w, method)
		}
	}
}
