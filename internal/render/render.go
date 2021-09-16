package render

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var (
	app             *config.AppConfig
	pathToTemplates = "./templates"
	functions       = template.FuncMap{}
)

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td models.TemplateData) error {
	var tc map[string]*template.Template
	switch app.UseTemplateCache {
	case app.UseTemplateCache:
		tc = app.TemplateCache
	default:
		tc, _ = CreateTemplateCache()
	}
	t, ok := tc[tmpl]
	if !ok {
		return errors.New("could not get template from cache")
	}
	buf := new(bytes.Buffer)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println("could not render template from buffer")
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
		return err
	}
	return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {
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
