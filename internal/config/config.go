package config

import (
	"context"
	"database/sql"
	"html/template"
)

type AppConfig struct {
	TemplateCache  map[string]*template.Template
	Str            string
	DB             *sql.DB
	Session        string
	Ctx            context.Context
	ChanIdSupplier chan int
	ChanMutex      chan int
	TimeFormat     string
}
