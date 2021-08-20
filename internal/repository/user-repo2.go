package repository

import (
	"database/sql"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"log"
	"net/http"
)

type ControllerInterface interface {
	Create(u *model.User) http.HandlerFunc
	Get(email *string) http.HandlerFunc
	GetAll() http.HandlerFunc
	Delete(id int) http.HandlerFunc
	Update(id int, u *model.User) http.HandlerFunc

}

type Controllers struct {
	App     *config.AppConfig

}
const TableUser ="user"
var user model.User

func (c Controllers) Create(u *model.User, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sqlStmt := fmt.Sprintf("INSERT INTO %s (name, email, phone_number, password, status) VALUES(?,?,?,?,?) ", TableUser)

		p, err := db.Prepare(sqlStmt)
		CheckErr(err)
		res, err := p.Exec(u.Name, u.Email, u.PhoneNumber, u.Password, u.Status)
		CheckErr(err)
		id, err := res.LastInsertId()
		CheckErr(err)
		fmt.Println(id)

		fmt.Fprintf(w, u.Name, "this user name")
	}
}

func (c Controllers) Get(email *string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE email=?", TableUser)
		err := c.App.DB.QueryRow(sqlStmt, *email).Scan(&user.ID, &user.Name, &user.Email, &user.PhoneNumber, &user.Password, &user.Status)
		CheckErr(err)
		fmt.Println(user.Email)
	}
}

func (c Controllers) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []model.User
		sqlStmt := fmt.Sprintf("SELECT * FROM %s", TableUser)
		results, err := c.App.DB.Query(sqlStmt)
		CheckErr(err)
		for results.Next() {
			err = results.Scan(
				&user.ID,
				&user.Name,
				&user.Email,
				&user.PhoneNumber,
				&user.Password,
				&user.Status)
			CheckErr(err)
			users = append(users, user)
		}
	}
}

func (c Controllers) Delete(id int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", TableUser)
		_, err := c.App.DB.Exec(sqlStmt, id)
		CheckErr(err)
	}
}

func (c Controllers) Update(id int, u *model.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sqlStmt := fmt.Sprintf("UPDATE %s SET name=?, email=?, phone_number=?, password=?, status=? WHERE id=%d ", TableUser, id)
		stmt, err := c.App.DB.Prepare(sqlStmt)
		CheckErr(err)

		_, err = stmt.Exec(
			u.Name,
			u.Email,
			u.PhoneNumber,
			u.Password,
			u.Status)
		CheckErr(err)
	}
}


func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}