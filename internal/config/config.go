package config

import (
	"database/sql"
	"html/template"
	"log"
)

type AppConfig struct{
	TemplateCache map[string] *template.Template
	InfoLog *log.Logger
	ErrorLog *log.Logger
	DB *sql.DB


}
