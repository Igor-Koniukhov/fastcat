package controllers

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"net/http"
)

type CartControllerI interface {
	CreateCart(db *sql.DB) http.HandlerFunc
	GetCart(db *sql.DB) http.HandlerFunc
	GetAllCarts(db *sql.DB) http.HandlerFunc
	DeleteCart(db *sql.DB) http.HandlerFunc
	UpdateCart(db *sql.DB) http.HandlerFunc
}

type CartControllers struct {
	App *config.AppConfig
}

func (c CartControllers) CreateCart(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (c CartControllers) GetCart(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (c CartControllers) GetAllCarts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (c CartControllers) DeleteCart(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (c CartControllers) UpdateCart(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

