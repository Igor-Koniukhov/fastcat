package config

import (
	"database/sql"
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

type AppConfig struct {
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	WarningLog    *log.Logger
	Session       *scs.SessionManager
	AccessToken   string
	RefreshToken  string
	Str           string
	DB            *sql.DB
}
