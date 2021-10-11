package handlers

import (
	"database/sql"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/handlers/repoForTest"
	"github.com/igor-koniukhov/fastcat/internal/render"

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

type TestRepository struct{
	repoForTest.UserTestRepository
	repoForTest.SupplierTestRepository
	repoForTest.ProductTestRepository
	repoForTest.OrderTestRepository
	repoForTest.CartTestRepository
}

func NewTestRepository(app *config.AppConfig, db *sql.DB) *TestRepository {
	return &TestRepository{
		UserTestRepository:     repoForTest.NewUserTestRepository(app, db),
		SupplierTestRepository: repoForTest.NewSupplierTestRepository(app, db),
		ProductTestRepository:  repoForTest.NewProductTestRepository(app, db),
		OrderTestRepository:    repoForTest.NewOrderTestRepository(app, db),
		CartTestRepository:     repoForTest.NewCartTestRepository(app, db),
	}
}

func NewTestHandlers(app *config.AppConfig, repos *TestRepository) *Handlers {
	return &Handlers{
		User:     NewUserHandler(app, repos.UserTestRepository),
		Supplier: NewSupplierHandler(app, repos.SupplierTestRepository),
		Product:  NewProductHandler(app, repos.ProductTestRepository),
		Order:    NewOrderHandler(app, repos.OrderTestRepository),
		Cart:     NewCartHandler(app, repos.CartTestRepository),
	}
}

func getRoutes() http.Handler {
	render.NewTemplates(&app)
	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("could not parse template cache")

	}
	app.TemplateCache = tc
	app.UseTemplateCache = true

	repo := NewTestRepository(&app, db)
	www := NewTestHandlers(&app, repo)


	mux := http.NewServeMux()
	mux.HandleFunc("/login", www.User.PostLogin)
	mux.HandleFunc("/show-login", www.User.ShowLogin)
	mux.HandleFunc("/refreshToken", www.User.Refresh)
	mux.HandleFunc("/logout", www.User.Logout)

	mux.HandleFunc("/registration", www.User.ShowRegistration)
	mux.HandleFunc("/user/create", www.User.SingUp)
	mux.HandleFunc("/user/", www.User.Get)//check handler without Auth
	mux.HandleFunc("/users", www.User.GetAll)
	mux.HandleFunc("/user/update/", www.User.Update)//check handler without Auth
	mux.HandleFunc("/user/delete/",www.User.Delete)//check handler without Auth

	mux.HandleFunc("/order/create",www.Order.Create)//check handler without Auth
	mux.HandleFunc("/order/", www.Order.Get)//check handler without Auth
	mux.HandleFunc("/orders", www.Order.GetAll)//check handler without Auth
	mux.HandleFunc("/order/update/", www.Order.Update)//check handler without Auth
	mux.HandleFunc("/order/delete/", www.Order.Delete)//check handler without Auth

	mux.HandleFunc("/supplier/create",www.Supplier.Create)//check handler without Auth
	mux.HandleFunc("/supplier/", www.Supplier.Get)
	mux.HandleFunc("/suppliers", www.Supplier.GetAll)
	mux.HandleFunc("/supplier/update/", www.Supplier.Update)//check handler without Auth
	mux.HandleFunc("/supplier/delete/", www.Supplier.Delete)//check handler without Auth

	mux.HandleFunc("/product/create", www.Product.Create)//check handler without Auth
	mux.HandleFunc("/product/", www.Product.Get)
	mux.HandleFunc("/products", www.Product.GetAll)
	mux.HandleFunc("/product/update/", www.Product.Update)//check handler without Auth
	mux.HandleFunc("/product/delete/", www.Product.Delete)//check handler without Auth

	mux.HandleFunc("/cart/create", www.Cart.Create)
	mux.HandleFunc("/cart/", www.Cart.Get)
	mux.HandleFunc("/carts", www.Cart.GetAll)


	fileServe := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServe))

	return mux
}
func getAuthRoutes() http.Handler {
	render.NewTemplates(&app)
	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("could not parse template cache")

	}
	app.TemplateCache = tc
	app.UseTemplateCache = true

	repo := NewTestRepository(&app, db)
	www := NewTestHandlers(&app, repo)


	mux := http.NewServeMux()


	mux.Handle("/user/", Auth(http.HandlerFunc(www.User.Get)))
	mux.Handle("/user/update/", Auth(http.HandlerFunc(www.User.Update)))
	mux.Handle("/user/delete/", Auth(http.HandlerFunc(www.User.Delete)))

	mux.Handle("/order/create", Auth(http.HandlerFunc(www.Order.Create)))
	mux.Handle("/order/", Auth(http.HandlerFunc(www.Order.Get)))
	mux.Handle("/orders", Auth(http.HandlerFunc(www.Order.GetAll)))
	mux.Handle("/order/update/", Auth(http.HandlerFunc(www.Order.Update)))
	mux.Handle("/order/delete/", Auth(http.HandlerFunc(www.Order.Delete)))

	mux.Handle("/supplier/create", Auth(http.HandlerFunc(www.Supplier.Create)))
	mux.Handle("/supplier/update/", Auth(http.HandlerFunc(www.Supplier.Update)))
	mux.Handle("/supplier/delete/", Auth(http.HandlerFunc(www.Supplier.Delete)))

	mux.Handle("/product/create", Auth(http.HandlerFunc(www.Product.Update)))
	mux.Handle("/product/update/", Auth(http.HandlerFunc(www.Product.Update)))
	mux.Handle("/product/delete/", Auth(http.HandlerFunc(www.Product.Delete)))

	mux.Handle("/cart/update/", Auth(http.HandlerFunc(www.Cart.Update)))
	mux.Handle("/cart/delete/", Auth(http.HandlerFunc(www.Cart.Delete)))

	fileServe := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServe))

	return mux
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusSeeOther)
		//_, err := r.Cookie("Authorized")
		/*if err !=nil {
			http.Redirect(w, r, "/show-login", http.StatusSeeOther)
			return
		}*/
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
