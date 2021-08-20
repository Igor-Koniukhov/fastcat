package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/igor-koniukhov/fastcat/helpers"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/handlers"
	r "github.com/igor-koniukhov/fastcat/internal/repository"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
)

var (
	db *sql.DB
	err error
	app      config.AppConfig
	infoLog  *log.Logger
	errorLog *log.Logger
)

func init() {
	gotenv.Load()
}

func main() {
	port := os.Getenv("PORT")
	db, err = sql.Open("mysql", "fastcat:"+"PASSWORD"+"@tcp(127.0.0.1:"+"PORT_DB"+")/fastcat_db")
	if err != nil {
		log.Print(err)
		return
	} else {
		fmt.Println("you connected to DB")
	}
	defer db.Close()
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = infoLog

	userFileRepository := r.NewUserRepository(&app)



	userCreateHandler := handlers.UserHandler{
		UserRepository: userFileRepository,
	}
	err = userCreateHandler.MyHandle()
	if err != nil {
		log.Panic(err)
	}
	helpers.NewHelpers(&app)

	http.Handle("/", http.FileServer(http.Dir("./public")))

	log.Fatal(http.ListenAndServe(port, nil))
}
