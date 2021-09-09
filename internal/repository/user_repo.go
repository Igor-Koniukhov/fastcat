package repository

import (
	"fmt"
	dr "github.com/igor-koniukhov/fastcat/driver"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	web "github.com/igor-koniukhov/webLogger/v3"
)

type UserRepository interface {
	Create(u *model.User) (*model.User, error)
	Get(id int) *model.User
	GetAll() []model.User
	Delete(id int) error
	Update(id int, u *model.User) *model.User
}

type UserRepo struct{
	App *config.AppConfig
}

func NewUserRepository(app *config.AppConfig) *UserRepo {
	return &UserRepo{App: app}
}

func (usr UserRepo) Create(u *model.User) (*model.User, error) {
	sqlStmt := fmt.Sprintf("INSERT INTO %s (name, email, phone_number, password, status) VALUES(?,?,?,?,?) ", dr.TableUser)
	p, err := usr.App.DB.Prepare(sqlStmt)
	defer p.Close()
	web.Log.Error(err, err)
	_, err = p.Exec(u.Name, u.Email, u.PhoneNumber, u.Password, u.Status)
	web.Log.Error(err, err)
	return u, err
}

func (usr UserRepo) Get(id int) *model.User{
	var user model.User
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id = ? ", dr.TableUser)
	err := usr.App.DB.QueryRow(sqlStmt, id).Scan(
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

func (usr UserRepo) GetAll() []model.User {
	var user model.User
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
	return users
}

func (usr UserRepo) Delete(id int) error {
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", dr.TableUser)
	_, err := usr.App.DB.Exec(sqlStmt, id)
	web.Log.Error(err, err)
	return err
}

func (usr UserRepo) Update(id int, u *model.User) *model.User {
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








