package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	web "github.com/igor-koniukhov/webLogger/v3"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	GetUserByID(id int) (*model.User, error)
	GetAll() []model.User
	Delete(id int) error
	Update(id int, user *model.User) *model.User
	GetUserByEmail(email string) (*model.User, error)

}

type UserRepo struct {
	App *config.AppConfig
	DB *sql.DB
	Users []model.User
	User  model.User
}

func NewUserRepository(app *config.AppConfig, DB *sql.DB) *UserRepo {
	return &UserRepo{App: app, DB: DB}
}

func (usr UserRepo) Create(user *model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("INSERT INTO %s (name, email, password) VALUES(?,?,?) ", model.TableUser)
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	web.Log.Error(err, err)
	_, err = usr.DB.ExecContext(ctx, sqlStmt,
		user.Name,
		user.Email,
		pass)
	web.Log.Error(err, err)
	return user, nil
}

func (usr UserRepo) GetUserByID(id int) (*model.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id = ? ", model.TableUser)
	row := usr.DB.QueryRowContext(ctx, sqlStmt, id)
	err :=row.Scan(
		&usr.User.ID,
		&usr.User.Name,
		&usr.User.Email,
		&usr.User.Password,
		&usr.User.DeletedAt,
		&usr.User.CreatedAT,
		&usr.User.UpdatedAT)
	web.Log.Error(err, err)
	return &usr.User, nil
}

func (usr UserRepo) GetAll() []model.User {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT * FROM %s", model.TableUser)
	results, err := usr.DB.QueryContext(ctx, sqlStmt)
	web.Log.Error(err, err)
	for results.Next() {
		err = results.Scan(
			&usr.User.ID,
			&usr.User.Name,
			&usr.User.Email,
			&usr.User.Password,
			&usr.User.DeletedAt,
			&usr.User.CreatedAT,
			&usr.User.UpdatedAT)
		web.Log.Error(err, err)
		usr.Users = append(usr.Users, usr.User)
	}
	return usr.Users
}

func (usr UserRepo) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", model.TableUser)
	_, err := usr.DB.ExecContext(ctx, sqlStmt, id)
	web.Log.Error(err, err)
	return nil
}

func (usr UserRepo) Update(id int, user *model.User) *model.User {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	fmt.Println(id)
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, name=?, email=?, password=? WHERE id=? ", model.TableUser)
	_, err := usr.DB.ExecContext(ctx, sqlStmt,
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		id)
	web.Log.Error(err, err)
	return user
}

func (usr UserRepo) GetUserByEmail(email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE email = ? ", model.TableUser)
	row := usr.DB.QueryRowContext(ctx, sqlStmt, email)
	err :=row.Scan(
		&usr.User.ID,
		&usr.User.Name,
		&usr.User.Email,
		&usr.User.Password,
		&usr.User.DeletedAt,
		&usr.User.CreatedAT,
		&usr.User.UpdatedAT)
	web.Log.Error(err, "message: ", err)
	return &usr.User, nil
}


