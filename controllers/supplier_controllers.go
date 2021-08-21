package controllers

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"net/http"
)

type SupplierControllerI interface {
	CreateSupplier(db *sql.DB, method string) http.HandlerFunc
	GetSupplier(db *sql.DB, method string) http.HandlerFunc
	GetAllSuppliers(db *sql.DB, method string) http.HandlerFunc
	DeleteSupplier(db *sql.DB, method string) http.HandlerFunc
	UpdateSupplier(db *sql.DB, method string) http.HandlerFunc
}

type SupplierControllers struct {
	App *config.AppConfig
}

func (s SupplierControllers) CreateSupplier(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (s SupplierControllers) GetSupplier(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

func (s SupplierControllers) GetAllSuppliers(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

func (s SupplierControllers) DeleteSupplier(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

func (s SupplierControllers) UpdateSupplier(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

