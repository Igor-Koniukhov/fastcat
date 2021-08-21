package controllers

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"net/http"
)

type SupplierControllerI interface {
	CreateSupplier(db *sql.DB) http.HandlerFunc
	GetSupplier(db *sql.DB) http.HandlerFunc
	GetAllSuppliers(db *sql.DB) http.HandlerFunc
	DeleteSupplier(db *sql.DB) http.HandlerFunc
	UpdateSupplier(db *sql.DB) http.HandlerFunc
}

type SupplierControllers struct {
	App *config.AppConfig
}

func (s SupplierControllers) CreateSupplier(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (s SupplierControllers) GetSupplier(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (s SupplierControllers) GetAllSuppliers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (s SupplierControllers) DeleteSupplier(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (s SupplierControllers) UpdateSupplier(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

