package controllers

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"net/http"
)

type OrderControllerI interface {
	CreateOrder(db *sql.DB, method string) http.HandlerFunc
	GetOrder(db *sql.DB, method string) http.HandlerFunc
	GetAllOrders(db *sql.DB, method string) http.HandlerFunc
	DeleteOrder(db *sql.DB, method string) http.HandlerFunc
	UpdateOrder(db *sql.DB, method string) http.HandlerFunc
}

type OrderControllers struct {
	App *config.AppConfig
}

func (o OrderControllers) CreateOrder(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

func (o OrderControllers) GetOrder(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

func (o OrderControllers) GetAllOrders(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

func (o OrderControllers) DeleteOrder(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}

func (o OrderControllers) UpdateOrder(db *sql.DB, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:


		default:
			methodMassage(w, method)
		}
	}
}
