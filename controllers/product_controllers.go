package controllers

import (
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/repository"
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

func (p ProductControllers) CreateProduct(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:

		default:
			methodMessage(w, method)
		}
	}
}

func (p ProductControllers) GetProduct(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			productAppConfigProvider(p.App)
		default:
			methodMessage(w, method)
		}
	}
}

func (p ProductControllers) GetAllProducts(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			productAppConfigProvider(p.App)
		default:
			methodMessage(w, method)
		}
	}
}

func (p ProductControllers) DeleteProduct(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			productAppConfigProvider(p.App)
		default:
			methodMessage(w, method)
		}
	}
}

func (p ProductControllers) UpdateProduct(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			productAppConfigProvider(p.App)
		default:
			methodMessage(w, method)
		}
	}
}
