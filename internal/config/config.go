package config

import (
	"context"
	"html/template"
)

type AppConfig struct {
	TemplateCache  map[string]*template.Template
	Str            string
	Session        string
	Ctx            context.Context
	ChanIdSupplier chan int
	ChanMutex      chan int
	TimeFormat     string
}
