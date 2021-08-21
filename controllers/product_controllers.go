package controllers

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"net/http"
)

type ProductControllerI interface {
	CreateProduct(db *sql.DB, method string) http.HandlerFunc
	GetProduct(db *sql.DB, method string) http.HandlerFunc
	GetAllProducts(db *sql.DB, method string) http.HandlerFunc
	DeleteProduct(db *sql.DB, method string) http.HandlerFunc
	UpdateProduct(db *sql.DB, method string) http.HandlerFunc
}

type ProductControllers struct {
	App *config.AppConfig
}

func (p ProductControllers) CreateProduct(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

func (p ProductControllers) GetProduct(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

func (p ProductControllers) GetAllProducts(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

func (p ProductControllers) DeleteProduct(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

func (p ProductControllers) UpdateProduct(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

