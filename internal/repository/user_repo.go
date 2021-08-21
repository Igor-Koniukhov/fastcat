package repository

import (
	"database/sql"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type UserRepositoryI interface {
	CreateUser(u *model.User, db *sql.DB) (*model.User, error)
	GetUser(nameParam, param *string, db *sql.DB) *model.User
	GetAllUsers(db *sql.DB) *[]model.User
	DeleteUser(id int, db *sql.DB) error
	UpdateUser(id int, u *model.User, db *sql.DB) *model.User
	Param(r *http.Request) (string, string, int)
}



type UserRepository struct {

}


const TableUser = "user"

func (usr *UserRepository) Param(r *http.Request) (string, string, int) {

	var paramName string
	var param string
	var id int
	fields := strings.Split(r.URL.String(), "/")
	str := fields[len(fields)-1]
	//TODO: need to be rewriting with regexp
	if len(str) > 5 {
		paramName = "email"
		param = str
		id = 0
	} else {
		num, err := strconv.Atoi(str)
		CheckErr(err)
		paramName = "id"
		param = str
		id = num
	}
	return param, paramName, id
}

func (usr *UserRepository) CreateUser(u *model.User, db *sql.DB) (*model.User, error) {
	sqlStmt := fmt.Sprintf("INSERT INTO %s (name, email, phone_number, password, status) VALUES(?,?,?,?,?) ", TableUser)
	p, err := db.Prepare(sqlStmt)
	defer p.Close()
	CheckErr(err)
	_, err = p.Exec(u.Name, u.Email, u.PhoneNumber, u.Password, u.Status)
	CheckErr(err)
	return u, err
}

func (usr UserRepository) GetUser(nameParam, param *string, db *sql.DB) *model.User {

	var user model.User
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE %s=?", TableUser, *nameParam)
	err := db.QueryRow(sqlStmt, *param).Scan(&user.ID, &user.Name, &user.Email, &user.PhoneNumber, &user.Password, &user.Status)
	CheckErr(err)
	return &user
}

var user model.User

func (usr *UserRepository) GetAllUsers(db *sql.DB) *[]model.User {
	var users []model.User
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
	return &users
}

func (usr *UserRepository) DeleteUser(id int, db *sql.DB) error {
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", TableUser)
	_, err := db.Exec(sqlStmt, id)
	CheckErr(err)
	return err
}

func (usr *UserRepository) UpdateUser(id int, u *model.User, db *sql.DB) *model.User {

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

	return u
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
