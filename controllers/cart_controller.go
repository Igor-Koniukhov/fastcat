package controllers

import (
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	"net/http"
)

type Cart interface {
	Create() http.HandlerFunc
	Get() http.HandlerFunc
	GetAll() http.HandlerFunc
	Delete() http.HandlerFunc
	Update() http.HandlerFunc
}

type CartController struct {
	repo dbrepo.CartRepository
}

func NewCartController(repo dbrepo.CartRepository) *CartController {
	return &CartController{repo: repo}
}

func (c CartController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (c CartController) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
	}
}

func (c CartController) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
	}
}

func (c CartController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (c CartController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
