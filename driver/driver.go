package driver

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

//DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}
var dbConn = &DB{}
const maxOpenDBConn = 10
const maxIdleDBConn = 5
const maxDBLifeTime = 5 * time.Minute

// ConnectSQL creates database pool for SQL
func ConnectDB(DSN string) (*DB, error) {
	d, err := NewDatabase("mysql", DSN)
	if err != nil {
		panic(err)
	}
	d.SetMaxOpenConns(maxOpenDBConn)
	d.SetMaxIdleConns(maxIdleDBConn)
	d.SetConnMaxLifetime(maxDBLifeTime)

	dbConn.SQL = d
	err = testDB(d)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}
// testDB ping to the DB
func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}
// NewDatabase creates new DB
func NewDatabase(d, dsn string) (*sql.DB, error) {
	DSN := os.Getenv(dsn)
	db, err := sql.Open(d, DSN)
	if err != nil {
		return nil, err
	} else {
		fmt.Println("you connected to DB")
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
