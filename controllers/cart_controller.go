package controllers

import (

	"github.com/igor-koniukhov/fastcat/internal/repository"
	"net/http"
)

type Cart interface {
	Create(method string) http.HandlerFunc
	Get(method string) http.HandlerFunc
	GetAll(method string) http.HandlerFunc
	Delete(method string) http.HandlerFunc
	Update(method string) http.HandlerFunc
}

type CartController struct {
	repo repository.CartRepositoryInterface
}

func NewCartController(repo repository.CartRepositoryInterface) *CartController {
	return &CartController{repo: repo}
}

func (c CartController) Create(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {


	}
}

func (c CartController) Get(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")


	}
}

func (c CartController) GetAll(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")

	}
}

func (c CartController) Delete(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (c CartController) Update(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {


	}
}
