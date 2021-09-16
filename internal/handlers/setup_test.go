package handlers

import (
	"database/sql"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/render"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"log"
	"path/filepath"

	"html/template"
	"net/http"
)

var (
	app             config.AppConfig
	db              *sql.DB
	pathToTemplates = "./../../templates"
	functions       = template.FuncMap{}
)

func getRoutes() http.Handler {
	render.NewTemplates(&app)
	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("could not parse template cache")

	}
	app.TemplateCache = tc
	app.UseTemplateCache = false

	repo := repository.NewRepository(&app, db)
	www := NewHandlers(&app, repo)
	repository.NewRepo(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/login", www.User.PostLogin)
	mux.HandleFunc("/show-login", www.User.ShowLogin)
	mux.HandleFunc("/refresh", www.User.Refresh)
	mux.HandleFunc("/logout", www.User.Logout)

	mux.HandleFunc("/registration", www.User.ShowRegistration)
	mux.HandleFunc("/user/create", www.User.Create)
	mux.Handle("/user/", Auth(http.HandlerFunc(www.User.Get)))
	mux.HandleFunc("/users", www.User.GetAll)
	mux.Handle("/user/update/", Auth(http.HandlerFunc(www.User.Update)))
	mux.Handle("/user/delete/", Auth(http.HandlerFunc(www.User.Delete)))

	mux.Handle("/order/create", Auth(http.HandlerFunc(www.Order.Create)))
	mux.Handle("/order/", Auth(http.HandlerFunc(www.Order.Get)))
	mux.Handle("/orders", Auth(http.HandlerFunc(www.Order.GetAll)))
	mux.Handle("/order/update/", Auth(http.HandlerFunc(www.Order.Update)))

	mux.Handle("/supplier/create", Auth(http.HandlerFunc(www.Supplier.Create)))
	mux.HandleFunc("/supplier/", www.Supplier.Get)
	mux.HandleFunc("/suppliers", www.Supplier.GetAll)
	mux.Handle("/supplier/update/", Auth(http.HandlerFunc(www.Supplier.Update)))
	mux.Handle("/supplier/delete/", Auth(http.HandlerFunc(www.Supplier.Delete)))

	mux.Handle("/product/create", Auth(http.HandlerFunc(www.Product.Update)))
	mux.HandleFunc("/product/", www.Product.Get)
	mux.HandleFunc("/products", www.Product.GetAll)
	mux.Handle("/product/update/", Auth(http.HandlerFunc(www.Product.Update)))
	mux.Handle("/product/delete/", Auth(http.HandlerFunc(www.Product.Delete)))

	mux.HandleFunc("/cart/create", www.Cart.Create)
	mux.HandleFunc("/cart/", www.Cart.Get)
	mux.HandleFunc("/carts", www.Cart.GetAll)
	mux.Handle("/cart/update/", Auth(http.HandlerFunc(www.Cart.Update)))
	mux.Handle("/cart/delete/", Auth(http.HandlerFunc(www.Cart.Delete)))

	fileServe := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServe))

	return mux
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("Authorized")
		if err !=nil {
			http.Redirect(w, r, "/show-login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func CreateTestTemplateCache() (map[string]*template.Template, error) {
	tCash := map[string]*template.Template{}
	pages, err := filepath.Glob(fmt.Sprintf("%v/*page.tmpl", pathToTemplates))
	if err != nil {
		return tCash, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return tCash, err
		}
		matches, err := filepath.Glob(fmt.Sprintf("%v/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return tCash, err
		}
		if len(matches) > 0 {
			ts, err := ts.ParseGlob(fmt.Sprintf("%v/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return tCash, err
			}
			tCash[name] = ts
		}
	}
	return tCash, nil
}
