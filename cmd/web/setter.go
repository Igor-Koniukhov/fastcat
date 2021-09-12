package main

import (
	"github.com/igor-koniukhov/fastcat/driver"
	web "github.com/igor-koniukhov/webLogger/v3"
	"time"
)


func SetAndRun()(*driver.DB, error)  {
	db, err := driver.ConnectDB("DSN")

	web.Log.Fatal(err, "Cannot connect to database!", err)
	app.Session = "This is session"
	app.TimeFormat = time.Now().UTC().Format("2006-01-02 15:04:05.999999")

	logSet := web.NewLogStruct(&web.LogParameters{
		OutWriter:  web.ConsoleAndFile,
		FilePath:   "./logs",
		LogFile:    "/logger.log",
		TimeFormat: "[15:04:05||2006.01.02]",
	})
	web.NewLog(logSet)
	return db, nil
}


