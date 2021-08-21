package controllers

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"net/http"
)

type CartControllerI interface {
	CreateCart(db *sql.DB, method string) http.HandlerFunc
	GetCart(db *sql.DB, method string) http.HandlerFunc
	GetAllCarts(db *sql.DB, method string) http.HandlerFunc
	DeleteCart(db *sql.DB, method string) http.HandlerFunc
	UpdateCart(db *sql.DB, method string) http.HandlerFunc
}

type CartControllers struct {
	App *config.AppConfig
}

func (c CartControllers) CreateCart(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

func (c CartControllers) GetCart(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

func (c CartControllers) GetAllCarts(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

func (c CartControllers) DeleteCart(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

func (c CartControllers) UpdateCart(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

