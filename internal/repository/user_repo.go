package repository

import (
	"database/sql"
	"fmt"
	dr "github.com/igor-koniukhov/fastcat/driver"
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
var user model.User

type UserRepository struct {}

func (usr *UserRepository) CreateUser(u *model.User, db *sql.DB) (*model.User, error) {
	sqlStmt := fmt.Sprintf("INSERT INTO %s (name, email, phone_number, password, status) VALUES(?,?,?,?,?) ", dr.TableUser)
	p, err := db.Prepare(sqlStmt)
	defer p.Close()
	CheckErr(err)
	_, err = p.Exec(u.Name, u.Email, u.PhoneNumber, u.Password, u.Status)
	CheckErr(err)
	return u, err
}

func (usr *UserRepository) GetUser(nameParam, param *string, db *sql.DB) *model.User {
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE %s = '%s' ", dr.TableUser, *nameParam, *param)
	err := db.QueryRow(sqlStmt).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PhoneNumber,
		&user.Password,
		&user.Status,
		&user.CreatedAT,
		&user.UpdatedAT)
	CheckErr(err)
	return &user
}

func (usr *UserRepository) GetAllUsers(db *sql.DB) *[]model.User {
	var users []model.User
	sqlStmt := fmt.Sprintf("SELECT * FROM %s", dr.TableUser)
	results, err := db.Query(sqlStmt)
	CheckErr(err)
	for results.Next() {
		err = results.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.PhoneNumber,
			&user.Password,
			&user.Status,
			&user.CreatedAT,
			&user.UpdatedAT)
		CheckErr(err)
		users = append(users, user)
	}
	return &users
}

func (usr *UserRepository) DeleteUser(id int, db *sql.DB) error {
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", dr.TableUser)
	_, err := db.Exec(sqlStmt, id)
	CheckErr(err)
	return err
}

func (usr *UserRepository) UpdateUser(id int, u *model.User, db *sql.DB) *model.User {

	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, name=?, email=?, phone_number=?, password=?, status=? WHERE id=%d ", dr.TableUser, id)
	stmt, err := db.Prepare(sqlStmt)
	CheckErr(err)
	_, err = stmt.Exec(
		u.ID,
		u.Name,
		u.Email,
		u.PhoneNumber,
		u.Password,
		u.Status)
	CheckErr(err)
	fmt.Println(*u)
	return u
}

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

func CheckErr(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
