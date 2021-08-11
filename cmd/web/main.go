package main

import (
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/handlers"
	r "github.com/igor-koniukhov/fastcat/internal/repository"
	"log"
	"net/http"
)

var app config.AppConfig

func main() {

	userFileRepository:=  r.NewUserRepository(&app)
	userCreateHandler := handlers.UserHandler{
		UserRepository: userFileRepository,
	}
	err := userCreateHandler.Handle()
	if err != nil {
		log.Panic(err)
	}

	http.Handle("/",  http.FileServer(http.Dir("./public")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
