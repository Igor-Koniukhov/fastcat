package dbrepo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	web "github.com/igor-koniukhov/webLogger/v2"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, int, error)
	SetUserSession(id int, token *models.LoginResponse) error
	GetUserSession(email string ) (*models.LoginResponse, error )
	UpdateSetUserSession(id int, token *models.LoginResponse) error
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
	App   *config.AppConfig
}

func NewUserRepository(app *config.AppConfig, DB *sql.DB) *UserRepo {
	return &UserRepo{App: app, DB: DB}
}

func (usr UserRepo) Create(user *models.User) (*models.User, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("INSERT INTO %s (name, email, tel, password) VALUES(?,?,?,?) ", models.TableUsers)
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		web.Log.Fatal(err)
		return nil, 0, err
	}
	res, err := usr.DB.ExecContext(ctx, sqlStmt,
		user.Name,
		user.Email,
		user.Tel,
		pass)
	if err != nil {
		web.Log.Fatal(err)
		return nil, 0, err
	}
	userId, err := res.LastInsertId()
	if err != nil {
		web.Log.Fatal(err)
		return nil, 0, err
	}
	return user, int(userId), nil
}

func (usr UserRepo) SetUserSession(id int, token *models.LoginResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmtSession := fmt.Sprintf("INSERT INTO %s (users_id, session) VALUES(?, ?) ", models.TableSessions)
	tokenBearer := models.LoginResponse{
		AccessToken:  "Bearer" + token.AccessToken,
		RefreshToken: "Bearer" + token.RefreshToken,
	}
	tok, err := json.Marshal(&tokenBearer)
	if err != nil {
		web.Log.Error(err)
		return err
	}
	_, err = usr.DB.ExecContext(ctx, sqlStmtSession,
		id,
		string(tok),
	)
	if err != nil {
		web.Log.Fatal(err)
		return err
	}
	return nil
}
func (usr UserRepo) UpdateSetUserSession(id int, token *models.LoginResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmtSession := fmt.Sprintf("UPDATE %s SET session=? WHERE users_id=? ", models.TableSessions)
	tokenBearer := models.LoginResponse{
		AccessToken:  "Bearer" + token.AccessToken,
		RefreshToken: "Bearer" + token.RefreshToken,
	}
	tok, err := json.Marshal(&tokenBearer)
	if err != nil {
		web.Log.Error(err)
		return err
	}
	_, err = usr.DB.ExecContext(ctx, sqlStmtSession, string(tok), id)
	if err != nil {
		web.Log.Fatal(err)
		return err
	}
	return nil
}

func (usr UserRepo) GetUserSession(email string ) (*models.LoginResponse, error ){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	 user, err := usr.GetUserByEmail(email)
	 if err !=nil{
	 	web.Log.Error(err)
	 	return nil, err
	 }
	sqlStmtSession := fmt.Sprintf("SELECT session FROM %s  WHERE  users_id = ? ", models.TableSessions)
		var usersSessions models.UsersSessions
	row := usr.DB.QueryRowContext(ctx, sqlStmtSession, user.ID)
	var token *models.LoginResponse
	row.Scan(
		&usersSessions.Session,
		)
	err = json.Unmarshal([]byte(usersSessions.Session), &token)
		if err!=nil{
		web.Log.Error(err)
		return nil, err
	}
	return token, nil
}

func (usr UserRepo) GetUserByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id = ? ", models.TableUsers)
	row := usr.DB.QueryRowContext(ctx, sqlStmt, id)
	err := row.Scan(
		&usr.User.ID,
		&usr.User.Name,
		&usr.User.Email,
		&usr.User.Tel,
		&usr.User.Password,
		&usr.User.DeletedAt,
		&usr.User.CreatedAT,
		&usr.User.UpdatedAT)
	if err != nil {
		web.Log.Fatal(err)
		return nil, err
	}
	return &usr.User, nil
}

func (usr UserRepo) GetAll() []models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT * FROM %s", models.TableUsers)
	results, err := usr.DB.QueryContext(ctx, sqlStmt)
	if err != nil {
		web.Log.Fatal(err)
		return nil
	}
	for results.Next() {
		err = results.Scan(
			&usr.User.ID,
			&usr.User.Name,
			&usr.User.Email,
			&usr.User.Tel,
			&usr.User.Password,
			&usr.User.DeletedAt,
			&usr.User.CreatedAT,
			&usr.User.UpdatedAT)
		if err != nil {
			web.Log.Fatal(err)
			return nil
		}
		usr.Users = append(usr.Users, usr.User)
	}
	return usr.Users
}

func (usr UserRepo) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", models.TableUsers)
	_, err := usr.DB.ExecContext(ctx, sqlStmt, id)
	if err != nil {
		web.Log.Fatal(err)
		return err
	}
	return nil
}

func (usr UserRepo) Update(id int, user *models.User) *models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, name=?, email=?, tel=?, password=? WHERE id=? ", models.TableUsers)
	_, err := usr.DB.ExecContext(ctx, sqlStmt,
		user.ID,
		user.Name,
		user.Email,
		user.Tel,
		user.Password,
		id)
	if err != nil {
		web.Log.Fatal(err)
		return nil
	}
	return user
}

func (usr UserRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE email = ? ", models.TableUsers)
	row := usr.DB.QueryRowContext(ctx, sqlStmt, email)
	err := row.Scan(
		&usr.User.ID,
		&usr.User.Name,
		&usr.User.Email,
		&usr.User.Tel,
		&usr.User.Password,
		&usr.User.DeletedAt,
		&usr.User.CreatedAT,
		&usr.User.UpdatedAT)
	if err != nil {
		web.Log.Error(err)
		return nil, err
	}
	return &usr.User, nil
}
