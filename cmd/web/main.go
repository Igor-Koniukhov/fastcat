package main

import (
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/handlers"
	r "github.com/igor-koniukhov/fastcat/internal/repository"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
)

var app config.AppConfig

func init()  {
	gotenv.Load()
}

func main() {
	port := os.Getenv("PORT")

	userFileRepository:=  r.NewUserRepository(&app)
	userCreateHandler := handlers.UserHandler{
		UserRepository: userFileRepository,
	}
	err := userCreateHandler.Handle()
	if err != nil {
		log.Panic(err)
	}

	http.Handle("/",  http.FileServer(http.Dir("./public")))

	log.Fatal(http.ListenAndServe(port, nil))
}
