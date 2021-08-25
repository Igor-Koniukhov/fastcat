package controllers

import (
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"net/http"
)

type SupplierControllerI interface {
	CreateSupplier(method string) http.HandlerFunc
	GetSupplier(method string) http.HandlerFunc
	GetAllSuppliers(method string) http.HandlerFunc
	DeleteSupplier(method string) http.HandlerFunc
	UpdateSupplier(method string) http.HandlerFunc
}

var RepoSupplier *SupplierControllers

type SupplierControllers struct {
	App *config.AppConfig
}



func NewSupplierControllers(app *config.AppConfig) *SupplierControllers {
	return &SupplierControllers{App: app}
}

func NewControllersS(r *SupplierControllers) {
	RepoSupplier = r
}

func supplierAppConfigProvider(a *config.AppConfig) *repository.SupplierRepository {
	repo := repository.NewSupplierRepository(a)
	repository.NewRepoS(repo)
	return repo
}

func (s SupplierControllers) CreateSupplier(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			supplierAppConfigProvider(s.App)



		default:
			methodMessage(w, method)
		}
	}
}

func (s SupplierControllers) GetSupplier(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			supplierAppConfigProvider(s.App)

		default:
			methodMessage(w, method)
		}
	}
}

func (s SupplierControllers) GetAllSuppliers(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			supplierAppConfigProvider(s.App)

		default:
			methodMessage(w, method)
		}
	}
}

func (s SupplierControllers) DeleteSupplier(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			supplierAppConfigProvider(s.App)

		default:
			methodMessage(w, method)
		}
	}
}

func (s SupplierControllers) UpdateSupplier(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			supplierAppConfigProvider(s.App)

		default:
			methodMessage(w, method)
		}
	}
}
