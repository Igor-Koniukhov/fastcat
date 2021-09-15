package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	web "github.com/igor-koniukhov/webLogger/v3"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	GetAll() []models.User
	Delete(id int) error
	Update(id int, user *models.User) *models.User
	GetUserByEmail(email string) (*models.User, error)
}

type UserRepo struct {
	DB    *sql.DB
	Users []models.User
	User  models.User
	App *config.AppConfig
}

func NewUserRepository(app *config.AppConfig, DB *sql.DB) *UserRepo {
	return &UserRepo{ App: app, DB: DB}
}

func (usr UserRepo) Create(user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("INSERT INTO %s (name, email, password) VALUES(?,?,?) ", models.TableUser)
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	_, err = usr.DB.ExecContext(ctx, sqlStmt,
		user.Name,
		user.Email,
		pass)
	if err != nil {
		log.Println(err)
	}
	return user, nil
}

func (usr UserRepo) GetUserByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id = ? ", models.TableUser)
	row := usr.DB.QueryRowContext(ctx, sqlStmt, id)
	err := row.Scan(
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

func (usr UserRepo) GetAll() []models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT * FROM %s", models.TableUser)
	results, err := usr.DB.QueryContext(ctx, sqlStmt)
	if err != nil {
		log.Println(err)
	}
	for results.Next() {
		err = results.Scan(
			&usr.User.ID,
			&usr.User.Name,
			&usr.User.Email,
			&usr.User.Password,
			&usr.User.DeletedAt,
			&usr.User.CreatedAT,
			&usr.User.UpdatedAT)
		if err != nil {
			log.Println(err)
		}
		usr.Users = append(usr.Users, usr.User)
	}
	return usr.Users
}

func (usr UserRepo) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", models.TableUser)
	_, err := usr.DB.ExecContext(ctx, sqlStmt, id)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (usr UserRepo) Update(id int, user *models.User) *models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, name=?, email=?, password=? WHERE id=? ", models.TableUser)
	_, err := usr.DB.ExecContext(ctx, sqlStmt,
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		id)
	if err != nil {
		log.Println(err)
	}
	return user
}

func (usr UserRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE email = ? ", models.TableUser)
	row := usr.DB.QueryRowContext(ctx, sqlStmt, email)
	err := row.Scan(
		&usr.User.ID,
		&usr.User.Name,
		&usr.User.Email,
		&usr.User.Password,
		&usr.User.DeletedAt,
		&usr.User.CreatedAT,
		&usr.User.UpdatedAT)
	if err != nil {
		log.Println(err)
	}
	return &usr.User, nil
}


