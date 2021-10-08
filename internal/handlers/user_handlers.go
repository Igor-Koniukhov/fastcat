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
	"github.com/igor-koniukhov/fastcat/services/router"
	web "github.com/igor-koniukhov/webLogger/v2"
	"log"
	"net/http"
	"os"
)

type User interface {
	ShowRegistration(w http.ResponseWriter, r *http.Request)
	AboutUs(w http.ResponseWriter, r *http.Request)
	Contacts(w http.ResponseWriter, r *http.Request)
	SingUp(w http.ResponseWriter, r *http.Request)
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
func (us *UserHandler) ShowRegistration(w http.ResponseWriter, r *http.Request) {
	err := render.TemplateRender(w, r, "sign_up.page.tmpl", &models.TemplateData{})
	if err != nil {
		log.Fatal("cannot render template")
	}
	w.WriteHeader(http.StatusOK)
}
func (us *UserHandler) AboutUs(w http.ResponseWriter, r *http.Request) {
	err := render.TemplateRender(w, r, "about.page.tmpl", &models.TemplateData{})
	if err != nil {
		web.Log.Fatal(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (us *UserHandler) Contacts(w http.ResponseWriter, r *http.Request) {
	err := render.TemplateRender(w, r, "contacts.page.tmpl", &models.TemplateData{})
	if err != nil {
		web.Log.Fatal(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (us *UserHandler) SingUp(w http.ResponseWriter, r *http.Request) {
	var u models.User
		u.Name = r.FormValue("user-name")
		u.Email = r.FormValue("user-email")
		u.Password = r.FormValue("password")
	if us.checkUserExist(u.Email){
		web.Log.Fatal("user already exists")
		http.Redirect(w, r, "/registration", http.StatusSeeOther)
		return
	}
		_, err := us.repo.Create(&u)
		if err != nil {
			log.Println(err)
		}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}
func (us *UserHandler) ShowLogin(w http.ResponseWriter, r *http.Request) {
	err := render.TemplateRender(w, r, "show_login.page.tmpl", &models.TemplateData{})
	if err != nil {
		log.Fatal("cannot render template")
	}
	w.WriteHeader(http.StatusOK)
}
func (us *UserHandler) PostLogin(w http.ResponseWriter, r *http.Request) {
	setAuth := &http.Cookie{
		Name:  "Authorization",
		Value: us.App.BearerString,
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
	resp, _, err := services.TokenResponder(w, logReq)
	if err != nil {
		log.Println(err)
		r = r.WithContext(context.WithValue(r.Context(), "error", "Invalid login credentials"))
	}
	tokenString, err := services.GetTokenFromBearerString("Authorization" + setAuth.Value)
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
		Name:  user.Email,
		Value: resp.RefreshToken,
	}
	us.App.NameForGreet = user.Name
	http.SetCookie(w, auth)
	http.Redirect(w, r, "/products", http.StatusSeeOther)
	w.WriteHeader(http.StatusOK)
}

func (us *UserHandler) Refresh(w http.ResponseWriter, r *http.Request) {
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

func (us *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	if len(cookies) >= 0 {
		for _, ck := range cookies {
			if ck.Name == "Authorization" || ck.Name == "Bearer" {
				ck.MaxAge = -1
				http.SetCookie(w, ck)
			}
		}
		http.Redirect(w, r, "/show-login", http.StatusSeeOther)
		return
	}
}

func (us *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := router.GetKeyInt(r, ":id")
	user, err := us.repo.GetUserByID(id)
	if err != nil {
		log.Println(err)
	}
	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}
func (us *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users := us.repo.GetAll()
	err := json.NewEncoder(w).Encode(&users)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}
func (us *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := router.GetKeyInt(r, ":id")
	err := us.repo.Delete(id)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusAccepted)
}
func (us *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println(err)
	}
	id := router.GetKeyInt(r, ":id")
	user := us.repo.Update(id, &u)
	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}

func (us UserHandler)checkUserExist(email string) (ok bool)  {
	user, _ := us.repo.GetUserByEmail(email)
	if ok := email == user.Email; ok {
		web.Log.Info("users exists ", ok)
		return ok
	}
	return ok
}