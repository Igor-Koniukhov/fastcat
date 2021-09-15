package services

import (
	"github.com/igor-koniukhov/fastcat/driver"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/parser"
	web "github.com/igor-koniukhov/webLogger/v3"
	"log"
	"os"
	"time"
)
type Setter struct{
	App *config.AppConfig
}

func NewSetter(app *config.AppConfig) *Setter {
	return &Setter{App: app}
}

func(s *Setter) SetAndRun()(*driver.DB, error)  {
	db, err := driver.ConnectDB("DSN")
	rest := parser.NewRestMenuParser(s.App)
	parser.NewRestMenu(rest)
	if err !=nil {
		log.Fatal(err)
	}
	s.App.TimeFormat = time.Now().UTC().Format("2006-01-02 15:04:05.999999")
	s.App.BearerString = os.Getenv("BearerString")
	logSet := web.NewLogStruct(&web.LogParameters{
		OutWriter:  web.ConsoleAndFile,
		FilePath:   "./logs",
		LogFile:    "/logger.log",
		TimeFormat: "[15:04:05||2006.01.02]",
	})
	web.NewLog(logSet)
	return db, nil
}


