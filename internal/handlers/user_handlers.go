package handlers

import (
	"context"
	"encoding/json"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	"github.com/igor-koniukhov/fastcat/services"
	web "github.com/igor-koniukhov/webLogger/v3"
	"html/template"
	"log"
	"net/http"
)

type User interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	PostLogin(w http.ResponseWriter, r *http.Request)
	ShowLogin(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
	LogOut(w http.ResponseWriter, r *http.Request)
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
func (c *UserHandler) ShowLogin(w http.ResponseWriter, r *http.Request) {
	cookAuth := &http.Cookie{
		Name:  "Bearer",
		Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NCwiZXhwIjoxNjMxNDg0NzE0fQ.b0j7v7JWmpeJ5tV13nq2jXumTWLYIcO_lTZWjOrSwB8",
		Path:  "/login",
	}
		http.SetCookie(w, cookAuth)
	tpl, err := template.ParseGlob("public/*html")

	err = tpl.ExecuteTemplate(w, "show-login.html", cookAuth.Value)
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *UserHandler) PostLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	web.Log.Error(err)
	logReq := &models.LoginRequest{
		Email:    r.Form.Get("checkEmail"),
		Password: r.Form.Get("checkPass"),
	}
	_, id, err := services.TokenResponder(w, logReq)
	if err != nil {
		log.Println(err)
		r = r.WithContext(context.WithValue(r.Context(), "error", "Invalid login credentials"))
	}
	web.Log.Error(err, "message: ", err)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", id))
	r.WithContext(context.WithValue(r.Context(), "flash", "logged in successfully"))
}

func (c *UserHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	logReq := new(models.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&logReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, _, err := services.TokenResponder(w, logReq)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err)
	}
}

func (c *UserHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	cookAuth := &http.Cookie{
		Name:  "Bearer",
		Value: "",
		Path:  "/login",
	}
	http.SetCookie(w, cookAuth)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
