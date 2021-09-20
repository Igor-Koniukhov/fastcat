package main

import (
	"database/sql"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"net/http"
	"testing"
)

func TestRoutes(t *testing.T)  {
	var app config.AppConfig
	var db *sql.DB
	mux := routes(&app, db)
	switch v :=mux.(type) {
	case *http.ServeMux:
	// do nothing; test passed
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, type is %T", v))
	}
}