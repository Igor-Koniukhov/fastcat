package main

import (
	"github.com/alexedwards/scs/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/igor-koniukhov/fastcat/driver"
	"github.com/igor-koniukhov/fastcat/helpers"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/handlers"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
	"time"
)
const (
	Error = "\x1b[31;1m [ERROR] \033[0m"
	Warning = "\x1b[33;1m [WARNING] \033[0m"
	Info = "\x1b[34;1m [INFO] \033[0m"
)


var (
	app        config.AppConfig
	session *scs.SessionManager
	infoLog    *log.Logger
	errorLog   *log.Logger
	warningLog *log.Logger

)

func init() {
	gotenv.Load()
}

func main() {
	port := os.Getenv("PORT")
	DSN := os.Getenv("DSN")
	db := driver.ConnectMySQLDB(DSN)
	defer db.Close()

	app.DB = db
	infoLog = log.New(os.Stdout,Info, log.Ldate|log.Ltime|log.Lshortfile)
	app.InfoLog = infoLog
	errorLog = log.New(os.Stdout,Error, log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog
	warningLog = log.New(os.Stdout,Warning, log.Ldate|log.Ltime|log.Lshortfile)
	app.WarningLog = warningLog


	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	handlers.UserHandlers(&app)
	handlers.OrderHandlers(&app)
	handlers.SupplierHandlers(&app)
	handlers.ProductHandlers(&app)
	handlers.CartHandlers(&app)

	helpers.NewHelpers(&app)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Fatal(http.ListenAndServe(port, nil))
}
