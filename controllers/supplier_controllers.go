package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	web "github.com/igor-koniukhov/webLogger/v2"
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
		w.Header().Set("Content-Type", "json")
		switch r.Method {
		case method:
			repo := supplierAppConfigProvider(s.App)
			param, nameParam, _ := repo.Param(r)
			user := repository.RepoS.GetSupplier(&nameParam, &param)
			json.NewEncoder(w).Encode(&user)
		default:
			methodMessage(w, method)
		}
	}
}

func (s SupplierControllers) GetAllSuppliers(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		switch r.Method {
		case method:
			supplierAppConfigProvider(s.App)
			suppliers := repository.RepoS.GetAllUSuppliers()
			json.NewEncoder(w).Encode(&suppliers)
		default:
			methodMessage(w, method)
		}
	}
}

func (s SupplierControllers) DeleteSupplier(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			repo := supplierAppConfigProvider(s.App)
			_, _, id := repo.Param(r)
			err := repository.RepoS.DeleteSupplier(id)
			web.Log.Error(err, err)
			_, _ = fmt.Fprintf(w, fmt.Sprintf(" supplier with %d deleted", id))

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
