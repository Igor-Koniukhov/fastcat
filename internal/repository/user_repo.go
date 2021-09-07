package repository

import (
	"fmt"
	dr "github.com/igor-koniukhov/fastcat/driver"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	web "github.com/igor-koniukhov/webLogger/v3"
	"net/http"
	"strconv"
	"strings"
)
type UserRepositoryInterface interface {
	Create(u *model.User) (*model.User, error)
	Get(nameParam, param *string) *model.User
	GetAll() *[]model.User
	Delete(id int) error
	Update(id int, u *model.User) *model.User
	Param(r *http.Request) (string, string, int)
}

type UserRepository struct{
	App *config.AppConfig
}

func NewUserRepository(app *config.AppConfig) *UserRepository {
	return &UserRepository{App: app}
}

func (usr UserRepository) Create(u *model.User) (*model.User, error) {
	sqlStmt := fmt.Sprintf("INSERT INTO %s (name, email, phone_number, password, status) VALUES(?,?,?,?,?) ", dr.TableUser)
	p, err := usr.App.DB.Prepare(sqlStmt)
	defer p.Close()
	web.Log.Error(err, err)
	_, err = p.Exec(u.Name, u.Email, u.PhoneNumber, u.Password, u.Status)
	web.Log.Error(err, err)
	return u, err
}

func (usr UserRepository) Get(nameParam, param *string) *model.User {
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE %s = '%s' ", dr.TableUser, *nameParam, *param)
	err := usr.App.DB.QueryRow(sqlStmt).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PhoneNumber,
		&user.Password,
		&user.Status,
		&user.CreatedAT,
		&user.UpdatedAT)
	web.Log.Error(err, err)
	return &user
}

func (usr UserRepository) GetAll() *[]model.User {
	var users []model.User
	sqlStmt := fmt.Sprintf("SELECT * FROM %s", dr.TableUser)
	results, err := usr.App.DB.Query(sqlStmt)
	web.Log.Error(err, err)
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
		web.Log.Error(err, err)
		users = append(users, user)
	}
	return &users
}

func (usr UserRepository) Delete(id int) error {
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", dr.TableUser)
	_, err := usr.App.DB.Exec(sqlStmt, id)
	web.Log.Error(err, err)
	return err
}

func (usr UserRepository) Update(id int, u *model.User) *model.User {
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, name=?, email=?, phone_number=?, password=?, status=? WHERE id=%d ", dr.TableUser, id)

	stmt, err := usr.App.DB.Prepare(sqlStmt)

	web.Log.Error(err, err)
	_, err = stmt.Exec(
		u.ID,
		u.Name,
		u.Email,
		u.PhoneNumber,
		u.Password,
		u.Status)
	web.Log.Error(err, err)
	fmt.Println(*u)
	return u
}

func (usr UserRepository) Param(r *http.Request) (string, string, int) {
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
		web.Log.Error(err, err)
		paramName = "id"
		param = str
		id = num
	}
	return param, paramName, id
}

var user model.User






