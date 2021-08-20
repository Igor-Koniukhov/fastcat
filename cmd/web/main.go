package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/igor-koniukhov/fastcat/driver"
	"github.com/igor-koniukhov/fastcat/helpers"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
)

var (
	err      error
	app      config.AppConfig
	infoLog  *log.Logger
	errorLog *log.Logger
)

func init() {
	gotenv.Load()
}

var u model.User
var email = "arni@mail.com"
var id int = 21

func main() {
	port := os.Getenv("PORT")
	db := driver.ConnectDB(&app)
	defer db.Close()
	controller := repository.Controllers{}

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = infoLog
	/*repo := handlers.NewUserDBRepostitory(&app)
	handlers.NewHandlers(repo)*/

	http.HandleFunc("/user", controller.Create(db))
	http.HandleFunc("/user-email", controller.Get(&email, db))
	http.HandleFunc("/users", controller.GetAll(db))
	http.HandleFunc("/update", controller.Update(id, db))
	http.HandleFunc("/delete", controller.Delete(id, db))

	if err != nil {
		log.Panic(err)
	}
	helpers.NewHelpers(&app)

	http.Handle("/", http.FileServer(http.Dir("./public")))

	log.Fatal(http.ListenAndServe(port, nil))
}
