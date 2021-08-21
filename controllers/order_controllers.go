package controllers

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"net/http"
)

type OrderControllerI interface {
	CreateOrder(db *sql.DB) http.HandlerFunc
	GetOrder(db *sql.DB) http.HandlerFunc
	GetAllOrders(db *sql.DB) http.HandlerFunc
	DeleteOrder(db *sql.DB) http.HandlerFunc
	UpdateOrder(db *sql.DB) http.HandlerFunc
}

type OrderControllers struct {
	App *config.AppConfig
}

func (o OrderControllers) CreateOrder(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (o OrderControllers) GetOrder(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (o OrderControllers) GetAllOrders(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (o OrderControllers) DeleteOrder(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (o OrderControllers) UpdateOrder(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
