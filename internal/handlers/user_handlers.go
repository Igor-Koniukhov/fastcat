package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/render"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	"github.com/igor-koniukhov/fastcat/services"
	"github.com/igor-koniukhov/fastcat/services/router"
	web "github.com/igor-koniukhov/webLogger/v2"
	"log"
	"net/http"
)

type User interface {
	ShowRegistration(w http.ResponseWriter, r *http.Request)
	StatusPage(w http.ResponseWriter, r *http.Request)
	AboutUs(w http.ResponseWriter, r *http.Request)
	Contacts(w http.ResponseWriter, r *http.Request)
	SingUp(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	PostLogin(w http.ResponseWriter, r *http.Request)
	ShowLogin(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	Auth(next http.HandlerFunc) http.HandlerFunc
	RefreshToken(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	App  *config.AppConfig
	repo dbrepo.UserRepository
}

func NewUserHandler(app *config.AppConfig, repo dbrepo.UserRepository) *UserHandler {
	return &UserHandler{App: app, repo: repo}
}
func (us *UserHandler) StatusPage(w http.ResponseWriter, r *http.Request) {
	//Get data from context
	if auth := r.Context().Value("Authorization"); auth != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Authorized " + auth.(string) + "\n"))
		fmt.Println(auth.(string), "-Author")
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Logged in"))
	}
}
func (us *UserHandler) ShowRegistration(w http.ResponseWriter, r *http.Request) {
	err := render.TemplateRender(w, r, "sign_up.page.tmpl", &models.TemplateData{StringMap: us.App.TemplateInfo})
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
	mapForAutoFill := make(map[string]string)
	mapForAutoFill["Name"] = u.Name
	mapForAutoFill["Email"] = u.Email
	mapForAutoFill["Password"] = u.Password
	userCookie, err := r.Cookie("User")
	if err==nil{
		mapForAutoFill["UserName"] = userCookie.Value
	}
	us.App.TemplateInfo = mapForAutoFill
	if us.checkUserExists(u.Email) {
		mapForAutoFill["ErrorExistsUser"] = "User already exists!"
		http.Redirect(w, r, "/registration", http.StatusSeeOther)
		return
	}
	_, id, err := us.repo.Create(&u)
	if err != nil {
		log.Println(err)
	}
	token, err := services.TokenGenerator(w, id)
	if err != nil {
		web.Log.Error(err)
		return
	}
	setAccess := &http.Cookie{
		Name:       "Authorization",
		Value:      "Bearer "+token.AccessToken,
		HttpOnly:   true,
		SameSite:   0,
	}
	http.SetCookie(w, setAccess)
	err = us.repo.SetUserSession(id, token)
	if err != nil {
		web.Log.Error(err)
		return
	}
	setRefresh := &http.Cookie{
		Name:  "Refresh",
		Value: token.RefreshToken,
		HttpOnly:   true,
		SameSite:   0,
	}
	http.SetCookie(w, setRefresh)

	userGreet := &http.Cookie{
		Name:  "User",
		Value: u.Name,
	}
	http.SetCookie(w, userGreet)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}
func (us *UserHandler) ShowLogin(w http.ResponseWriter, r *http.Request) {
	_, err :=r.Cookie("Authorization")
	if err==nil  {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	err = render.TemplateRender(w, r, "show_login.page.tmpl", &models.TemplateData{
		ErrorMessage: us.App.ErrMessage,
		StringMap:    us.App.TemplateInfo})
	if err != nil {
		log.Fatal("cannot render template")
	}
	w.WriteHeader(http.StatusOK)
}
func (us *UserHandler) PostLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	logReq := &models.LoginRequest{
		Email:    r.Form.Get("checkEmail"),
		Password: r.Form.Get("checkPass"),
	}
	mapForAutoFill := make(map[string]string)
	mapForAutoFill["Email"] = logReq.Email
	mapForAutoFill["Password"] = logReq.Password
	us.App.TemplateInfo = mapForAutoFill

	if ok := us.checkUserExists(logReq.Email);!ok {
		us.App.ErrMessage = "User dose not exist"
		http.Redirect(w, r, "/show-login", http.StatusSeeOther)
		web.Log.Info(logReq.Email," ", ok)
		return
	}
	us.App.ErrMessage = ""
	token, id, err := services.TokenResponder(w, logReq)
	if err != nil {
		web.Log.Error(err)
		us.App.ErrMessage = "Invalid credentials"
		http.Redirect(w, r, "/show-login", http.StatusSeeOther)
		return
	}
	setAccess := &http.Cookie{
		Name:  "Authorization",
		Value: "Bearer "+token.RefreshToken,
		HttpOnly:   true,
		SameSite:   0,
	}
	http.SetCookie(w, setAccess)
	setRefresh := &http.Cookie{
		Name:  "Refresh",
		Value: token.RefreshToken,
		HttpOnly:   true,
		SameSite:   0,
	}
	http.SetCookie(w, setRefresh)
	us.App.ErrMessage = ""

	u, err := us.repo.GetUserByID(id)
	if err != nil {
		web.Log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userGreet := &http.Cookie{
		Name:  "User",
		Value: u.Name,
	}
	http.SetCookie(w, userGreet)
	http.Redirect(w, r, "/products", http.StatusSeeOther)
	w.WriteHeader(http.StatusOK)
}
func (us *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	if len(cookies) >= 0 {
		for _, ck := range cookies {
			if ck.Name == "Authorization" || ck.Name == "User" || ck.Name == "Refresh" {
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
func (us UserHandler) RefreshToken(w http.ResponseWriter, r *http.Request)  {
	us.App.ErrMessage = "refreshToken token"
	refresh, err :=r.Cookie("Refresh")
	if err != nil {
		web.Log.Error(err)
		return
	}
	auth, err :=r.Cookie("Authorization")
	if err != nil {
		web.Log.Error(err)
		return
	}
	auth.MaxAge=-1
	claims, err := services.ValidateToken(refresh.Value, services.RefreshSecret)
	if err != nil {
		web.Log.Error(err)
		return
	}
	refresh.MaxAge=-1
	token, err := services.TokenGenerator(w, claims.ID)
	if err != nil {
		web.Log.Error(err)
		return
	}
	user, err := us.repo.GetUserByID(claims.ID)
	if err != nil {
		web.Log.Error(err)
		return
	}
	setAccess := &http.Cookie{
		Name:  "Authorization",
		Value: "Bearer "+token.RefreshToken,
		HttpOnly:   true,
		SameSite:   0,
		MaxAge: 0,
	}
	http.SetCookie(w, setAccess)
	setRefresh := &http.Cookie{
		Name:  "Refresh",
		Value: token.RefreshToken,
		HttpOnly:   true,
		SameSite:   0,
	}
	http.SetCookie(w, setRefresh)

	userGreet := &http.Cookie{
		Name:  "User",
		Value: user.Name,
	}
	http.SetCookie(w, userGreet)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}
func (us UserHandler) checkUserExists(email string) (ok bool) {
	user, err := us.repo.GetUserByEmail(email)
	if err != nil {
		web.Log.Info(err, email, "-", ok)
		return ok
	}
	if ok := email == user.Email; ok {
		web.Log.Info(email," - User exists ", ok)
		return ok
	}
	return ok
}
