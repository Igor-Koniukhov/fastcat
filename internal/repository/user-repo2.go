package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"log"
	"net/http"
)

type ControllerInterface interface {
	Create(u *model.User, db *sql.DB) http.HandlerFunc
	Get(email *string, db *sql.DB) http.HandlerFunc
	GetAll(db *sql.DB) http.HandlerFunc
	Delete(id int, db *sql.DB) http.HandlerFunc
	Update(id int, u *model.User, db *sql.DB) http.HandlerFunc
}

type Controllers struct {
	App *config.AppConfig
}

const TableUser = "user"

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

func (c Controllers) Get(email *string, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE email=?", TableUser)
		err := db.QueryRow(sqlStmt, *email).Scan(&user.ID, &user.Name, &user.Email, &user.PhoneNumber, &user.Password, &user.Status)
		CheckErr(err)
		fmt.Println(user.Email)
	}
}

func (c Controllers) GetAll(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []model.User
		switch r.Method {
		case "GET":
			sqlStmt := fmt.Sprintf("SELECT * FROM %s", TableUser)
			results, err := db.Query(sqlStmt)
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
			json.NewEncoder(w).Encode(&users)
		default:
			http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)

		}
	}
}

func (c Controllers) Delete(id int, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "DELETE":
			sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", TableUser)
			_, err := db.Exec(sqlStmt, id)
			CheckErr(err)
			fmt.Fprintf(w, " user with "+string(id)+" deleted")
		default:
			http.Error(w, "Only DELETE method is allowed", http.StatusMethodNotAllowed)
		}
	}
}

func (c Controllers) Update(id int, u *model.User, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&u)
		switch r.Method {
		case "PUT":
			sqlStmt := fmt.Sprintf("UPDATE %s SET name=?, email=?, phone_number=?, password=?, status=? WHERE id=%d ", TableUser, id)
			stmt, err := db.Prepare(sqlStmt)
			CheckErr(err)

			_, err = stmt.Exec(
				u.Name,
				u.Email,
				u.PhoneNumber,
				u.Password,
				u.Status)
			CheckErr(err)
			json.NewEncoder(w).Encode(&u)
		default:

			http.Error(w, "Only UPDATE method is allowed", http.StatusMethodNotAllowed)
		}
	}
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
