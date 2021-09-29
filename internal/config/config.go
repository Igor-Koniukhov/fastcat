package config

import (
	"context"
	"html/template"
)

type AppConfig struct {
	TemplateCache    map[string]*template.Template
	UseTemplateCache bool
	Str              string
	Session          string
	BearerString     string
	UserAuthID       string
	NameForGreet     string
	Ctx              context.Context
	ChanIdSupplier   chan int
	ChanMutex        chan int
	TimeFormat       string
}
