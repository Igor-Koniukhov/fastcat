package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	web "github.com/igor-koniukhov/webLogger/v3"
	"net/http"
)

type User interface {
	Create() http.HandlerFunc
	Get() http.HandlerFunc
	GetAll() http.HandlerFunc
	Delete() http.HandlerFunc
	Update() http.HandlerFunc
}

type UserController struct {
	repo repository.UserRepository
}

func NewUserController(repo repository.UserRepository) *UserController {
	return &UserController{repo: repo}
}

func (c *UserController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u model.User
		json.NewDecoder(r.Body).Decode(&u)
		user, err := c.repo.Create(&u)
		web.Log.Error(err, err)
		json.NewEncoder(w).Encode(&user)
	}
}

func (c *UserController) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		id := param(r)
		user := c.repo.Get(id)
		json.NewEncoder(w).Encode(&user)
	}
}

func (c *UserController) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "json")
		users := c.repo.GetAll()
		json.NewEncoder(w).Encode(&users)
	}
}

func (c *UserController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := param(r)
		err := c.repo.Delete(id)
		web.Log.Error(err, err)
		_, _ = fmt.Fprintf(w, fmt.Sprintf(" user with %d deleted", id))

	}
}

func (c *UserController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u model.User
		json.NewDecoder(r.Body).Decode(&u)
		id := param(r)
		user := c.repo.Update(id, &u)
		json.NewEncoder(w).Encode(&user)
	}
}
