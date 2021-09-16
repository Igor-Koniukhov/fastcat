package handlers

import (
	"context"
	"encoding/json"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/render"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	"github.com/igor-koniukhov/fastcat/services"
	"log"
	"net/http"
	"os"
)

type User interface {
	ShowRegistration(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	PostLogin(w http.ResponseWriter, r *http.Request)
	ShowLogin(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	App  *config.AppConfig
	repo dbrepo.UserRepository
}

func NewUserHandler(app *config.AppConfig, repo dbrepo.UserRepository) *UserHandler {
	return &UserHandler{App: app, repo: repo}
}
func (c *UserHandler) ShowRegistration(w http.ResponseWriter, r *http.Request)  {
	err := render.RenderTemplate(w, r, "registration.page.tmpl", models.TemplateData{})
	if err !=nil {
		log.Fatal("cannot render template")
	}

}

func (c *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := r.ParseForm()
	u.Name = r.FormValue("user-name")
	u.Email = r.FormValue("user-email")
	u.Password = r.FormValue("password")
	user, err := c.repo.Create(&u)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
	http.Redirect(w,r, "/users", http.StatusSeeOther)
	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		log.Println(err)
	}
}

func (c *UserHandler) Get(w http.ResponseWriter, r *http.Request) {

	id := param(r)
	user, err := c.repo.GetUserByID(id)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		log.Println(err)
	}
}

func (c *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users := c.repo.GetAll()
	err := json.NewEncoder(w).Encode(&users)
	if err != nil {
		log.Println(err)
	}
}

func (c *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := param(r)
	err := c.repo.Delete(id)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusAccepted)
}

func (c *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var u models.User
	json.NewDecoder(r.Body).Decode(&u)
	id := param(r)
	user := c.repo.Update(id, &u)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&user)
	if err != nil {
		log.Println(err)
	}
}
func (c *UserHandler) ShowLogin(w http.ResponseWriter, r *http.Request) {
	err := render.RenderTemplate(w, r, "show_login.page.tmpl", models.TemplateData{})
	if err !=nil {
		log.Fatal("cannot render template")
	}
}

func (c *UserHandler) PostLogin(w http.ResponseWriter, r *http.Request) {
	setAuth := &http.Cookie{
		Name:  "Bearer",
		Value: c.App.BearerString,
	}
	http.SetCookie(w, setAuth)
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	logReq := &models.LoginRequest{
		Email:    r.Form.Get("checkEmail"),
		Password: r.Form.Get("checkPass"),
	}
	_, id, err := services.TokenResponder(w, logReq)
	if err != nil {
		log.Println(err)
		r = r.WithContext(context.WithValue(r.Context(), "error", "Invalid login credentials"))
	}
	tokenString, err := services.GetTokenFromBearerString("Bearer" + setAuth.Value)
	if err != nil {
		log.Println(err)
	}
	claims, err := services.ValidateToken(tokenString, os.Getenv("AccessSecret"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	user, err := repository.Repo.GetUserByID(claims.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	auth := &http.Cookie{
		Name:  "Authorized",
		Value: user.Email,
	}
	http.SetCookie(w, auth)
	http.Redirect(w, r, "/users", http.StatusSeeOther)
	w.WriteHeader(http.StatusOK)
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

func (c *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	if len(cookies) >= 0 {
		for _, ck := range cookies {
			if ck.Name == "Authorized" || ck.Name == "Bearer" {
				ck.MaxAge = -1
				http.SetCookie(w, ck)
			}
		}
		http.Redirect(w, r, "/show-login", http.StatusSeeOther)
	}
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
