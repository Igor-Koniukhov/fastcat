package handlers

import (
	"encoding/json"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	web "github.com/igor-koniukhov/webLogger/v3"
	"net/http"
)

type User interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	repo dbrepo.UserRepository
}

func NewUserHandler(repo dbrepo.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (c *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u models.User
	json.NewDecoder(r.Body).Decode(&u)
	user, err := c.repo.Create(&u)
	web.Log.Error(err, err)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&user)
	web.Log.Error(err)
}

func (c *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	id := param(r)
	user, err := c.repo.GetUserByID(id)
	web.Log.Error(err, "message: ", err)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&user)
	web.Log.Error(err)
}

func (c *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	users := c.repo.GetAll()
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&users)
	web.Log.Error(err)
}

func (c *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := param(r)
	err := c.repo.Delete(id)
	web.Log.Error(err, err)
	w.WriteHeader(http.StatusAccepted)
}

func (c *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var u models.User
	json.NewDecoder(r.Body).Decode(&u)
	id := param(r)
	user := c.repo.Update(id, &u)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&user)
	web.Log.Error(err)
}
