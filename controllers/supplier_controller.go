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

type SupplierControllerI interface {
	Create(method string) http.HandlerFunc
	Get(method string) http.HandlerFunc
	GetAll(method string) http.HandlerFunc
	Delete(method string) http.HandlerFunc
	Update(method string) http.HandlerFunc
}

var RepoSupplier *SupplierController

type SupplierController struct {
	App *config.AppConfig
}

func NewSupplierControllers(app *config.AppConfig) *SupplierController {
	return &SupplierController{App: app}
}

func NewControllersS(r *SupplierController) {
	RepoSupplier = r
}

func supplierAppConfigProvider(a *config.AppConfig) *repository.SupplierRepository {
	repo := repository.NewSupplierRepository(a)
	repository.NewRepoS(repo)
	return repo
}

func (s SupplierController) Create(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			var suppliers model.Suppliers
			json.NewDecoder(r.Body).Decode(&suppliers)
			_ = supplierAppConfigProvider(s.App)
			user, err := repository.RepoS.Create(&suppliers)
			checkError(err)
			json.NewEncoder(w).Encode(&user)
		default:
			methodMessage(w, method)
		}
	}
}

func (s SupplierController) Get(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		switch r.Method {
		case method:
			repo := supplierAppConfigProvider(s.App)
			param, nameParam, _ := repo.Param(r)
			user := repository.RepoS.Get(&nameParam, &param)
			json.NewEncoder(w).Encode(&user)
		default:
			methodMessage(w, method)
		}
	}
}

func (s SupplierController) GetAll(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		switch r.Method {
		case method:
			supplierAppConfigProvider(s.App)
			suppliers := repository.RepoS.GetAll()
			json.NewEncoder(w).Encode(&suppliers)

			methodMessage(w, method)
		}
	}
}

func (s SupplierController) Delete(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			repo := supplierAppConfigProvider(s.App)
			_, _, id := repo.Param(r)
			err := repository.RepoS.Delete(id)
			web.Log.Error(err, err)
			_, _ = fmt.Fprintf(w, fmt.Sprintf(" supplier with %d deleted", id))
		default:
			methodMessage(w, method)
		}
	}
}

func (s SupplierController) Update(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			var supplier model.Supplier
			json.NewDecoder(r.Body).Decode(&supplier)
			repo := supplierAppConfigProvider(s.App)
			_, _, id := repo.Param(r)
			user := repository.RepoS.Update(id, &supplier)
			json.NewEncoder(w).Encode(&user)
		default:
			methodMessage(w, method)
		}
	}
}
