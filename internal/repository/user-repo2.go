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
	Create( db *sql.DB) http.HandlerFunc
	Get(email *string, db *sql.DB) http.HandlerFunc
	GetAll(db *sql.DB) http.HandlerFunc
	Delete(id int, db *sql.DB) http.HandlerFunc
	Update(id int, db *sql.DB) http.HandlerFunc
}

type Controllers struct {
	App *config.AppConfig
}

const TableUser = "user"

var user model.User

func (c Controllers) Create( db *sql.DB) http.HandlerFunc {
	var u model.User
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			json.NewDecoder(r.Body).Decode(&u)
			sqlStmt := fmt.Sprintf("INSERT INTO %s (name, email, phone_number, password, status) VALUES(?,?,?,?,?) ", TableUser)
			p, err := db.Prepare(sqlStmt)
			CheckErr(err)
			_, err = p.Exec(u.Name, u.Email, u.PhoneNumber, u.Password, u.Status)
			CheckErr(err)
			json.NewEncoder(w).Encode(&u)
		default:
			methodMassage(w, "POST")
		}
	}
}

func (c Controllers) Get(email *string, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			var user model.User
			sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE email=?", TableUser)
			err := db.QueryRow(sqlStmt, *email).Scan(&user.ID, &user.Name, &user.Email, &user.PhoneNumber, &user.Password, &user.Status)
			CheckErr(err)
			fmt.Fprintf(w, user.Email)
		default:
			methodMassage(w, "GET")
		}
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
			methodMassage(w, "GET")
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
			methodMassage(w, "PUT")
		}
	}
}

func (c Controllers) Update(id int, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u model.User
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
			methodMassage(w, "DELETE")
		}
	}
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
func methodMassage(w http.ResponseWriter, m string)  {
	http.Error(w, "Only " +m+" method is allowed", http.StatusMethodNotAllowed)

}