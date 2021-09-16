package render

import (
	"fmt"
	"html/template"
	"path/filepath"
)
var (
	pathToTemplates = "./templates"
	functions = template.FuncMap{}
)



func CreateTemplateCache() (map[string]*template.Template, error)  {
	tCash := map[string]*template.Template{}
	pages, err := filepath.Glob(fmt.Sprintf("%v/*page.tmpl", pathToTemplates))
	if err!=nil{
			return tCash, err
	}
	for _, page := range pages{
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err !=nil {
			return tCash, err
		}
		matches, err := filepath.Glob(fmt.Sprintf("%v/*.layout.tmpl", pathToTemplates))
		if err !=nil {
			return tCash, err
		}
		if len(matches) >0 {
			ts, err := ts.ParseGlob(fmt.Sprintf("%v/*.layout.tmpl", pathToTemplates))
			if err !=nil {
				return tCash, err
			}
			tCash[name]=ts
		}
	}

	return tCash, nil
}
