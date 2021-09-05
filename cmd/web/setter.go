package main

import (
	"database/sql"


	"github.com/igor-koniukhov/webLogger/v2"
	"time"
)


func SetAppConfigParameters(db *sql.DB)  {
	app.DB = db
	app.Session = "This is session"
	app.TimeFormat = time.Now().UTC().Format("15:04:05||02.01.2006")

}


func SetWebLoggerParameters() {
	logSet := webLogger.NewLogStruct(&webLogger.LogParameters{
		OutWriter:  webLogger.ConsoleAndFile,
		FilePath:   "./logs",
		LogFile:    "/logger.log",
		TimeFormat: "[15:04:05||2006.01.02]",
	})
	webLogger.NewLog(logSet)
}
