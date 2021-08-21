package controllers

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"net/http"
)

type ProductControllerI interface {
	CreateProduct(db *sql.DB) http.HandlerFunc
	GetProduct(db *sql.DB) http.HandlerFunc
	GetAllProducts(db *sql.DB) http.HandlerFunc
	DeleteProduct(db *sql.DB) http.HandlerFunc
	UpdateProduct(db *sql.DB) http.HandlerFunc
}

type ProductControllers struct {
	App *config.AppConfig
}

func (p ProductControllers) CreateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (p ProductControllers) GetProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (p ProductControllers) GetAllProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (p ProductControllers) DeleteProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (p ProductControllers) UpdateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

