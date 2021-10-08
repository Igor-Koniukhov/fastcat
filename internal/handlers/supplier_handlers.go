package handlers

import (
	"encoding/json"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/render"
	"github.com/igor-koniukhov/fastcat/internal/repository/dbrepo"
	"github.com/igor-koniukhov/fastcat/services/router"
	web "github.com/igor-koniukhov/webLogger/v2"
	"log"
	"net/http"
	"strings"
)

type Supplier interface {
	Home(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetAllBySchedule(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type SupplierHandler struct {
	App  *config.AppConfig
	repo dbrepo.SupplierRepository
}

func NewSupplierHandler(app *config.AppConfig, repo dbrepo.SupplierRepository) *SupplierHandler {
	return &SupplierHandler{App: app, repo: repo}
}
func (s *SupplierHandler) Home(w http.ResponseWriter, r *http.Request) {
	suppliers := s.repo.GetAll()
	err := render.TemplateRender(w, r, "home.page.tmpl",
		&models.TemplateData{
			Suppliers: suppliers,
		})
	if err != nil {
		web.Log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
}
func (s *SupplierHandler) Create(w http.ResponseWriter, r *http.Request) {
	var suppliers models.Suppliers
	err := json.NewDecoder(r.Body).Decode(&suppliers)
	if err != nil {
		log.Println(err)
	}
	_, _, err = s.repo.Create(&suppliers)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
}
func (s *SupplierHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := router.GetKeyInt(r, ":id")
	supplier := s.repo.Get(id)
	w.WriteHeader(http.StatusOK)
	err := render.TemplateRender(w, r, "suppliersProducts.page.tmpl",
		&models.TemplateData{
			Supplier: supplier,
		})
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}
func (s *SupplierHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	suppliers := s.repo.GetAll()
	err := render.TemplateRender(w, r, "suppliers.page.tmpl",
		&models.TemplateData{
			Suppliers: suppliers,
		})
	if err != nil {
		web.Log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
}
func (s *SupplierHandler) GetAllBySchedule(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		web.Log.Fatal(err)
	}
	string := r.Form.Get("schedule")
	schedule := strings.Split(string, "--")
	start := schedule[0]
	end := schedule[1]
	suppliers := s.repo.GetAllBySchedule(start, end)
	err = render.TemplateRender(w, r, "suppliers.page.tmpl",
		&models.TemplateData{
			Suppliers: suppliers,
		})
	if err != nil {
		web.Log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
}
func (s *SupplierHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := router.GetKeyInt(r, ":id")
	err := s.repo.Delete(id)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusAccepted)
}
func (s *SupplierHandler) Update(w http.ResponseWriter, r *http.Request) {
	var supplier models.Supplier
	err := json.NewDecoder(r.Body).Decode(&supplier)
	if err != nil {
		log.Println(err)
	}
	id := router.GetKeyInt(r, ":id")
	user := s.repo.Update(id, &supplier)
	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}
