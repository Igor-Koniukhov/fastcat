package render

import (
	"github.com/igor-koniukhov/fastcat/internal/models"
	"net/http"
	"testing"
)

func TestTemplateRender(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
	var r *http.Request
	app.TemplateCache = tc

	var ww myWriter

	err = TemplateRender(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error writing template to browser", err)
	}

	err = TemplateRender(&ww, r, "non-existent.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("rendered template that does not exist")
	}

}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"

	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
