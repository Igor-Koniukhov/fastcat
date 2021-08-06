package main

import (
	"github.com/igor-koniukhov/fastcat/handlers"
	r "github.com/igor-koniukhov/fastcat/repository"
	"log"
)



func main() {

	userFileRepository:=  r.NewUserFileRepository()

	userCreateHandler := handlers.UserCreateHandler{
		UserRepository: userFileRepository,
	}
	err := userCreateHandler.Handle()
	if err != nil {
		log.Panic(err)
	}
}
