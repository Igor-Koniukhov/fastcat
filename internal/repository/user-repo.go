package repository

/*
import (
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"log"
)

type MethodRepositoryI interface {
	Create(u *model.User) (*model.User, error)
	Get(email *string) *model.User
	GetAll() *[]model.User
	Delete(id int) (*model.User, error)
	Update(id int, u *model.User) *model.User

}

type MethodRepo struct {
	App     *config.AppConfig

}

func NewUserDBRepository(a *config.AppConfig) *MethodRepo {
	return &MethodRepo{
		App:     a,

	}
}

const TableUser = "user"

func (usr MethodRepo) Create(u *model.User) (*model.User, error) {

	sqlStmt := fmt.Sprintf("INSERT INTO %s (name, email, phone_number, password, status) VALUES(?,?,?,?,?) ", TableUser)

	p, err := usr.App.DB.Prepare(sqlStmt)
	CheckErr(err)
	res, err := p.Exec(u.Name, u.Email, u.PhoneNumber, u.Password, u.Status)
	CheckErr(err)
	id, err := res.LastInsertId()
	CheckErr(err)
	fmt.Println(id)

	return u, err
}

func (usr MethodRepo) Get(email *string) *model.User {
	var user model.User
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE email=?", TableUser)
	err := usr.App.DB.QueryRow(sqlStmt, *email).Scan(&user.ID, &user.Name, &user.Email, &user.PhoneNumber, &user.Password, &user.Status)
	CheckErr(err)
	fmt.Println(user.Email)
	return &user

}

var user model.User

func (usr MethodRepo) GetAll() *[]model.User {
	var users []model.User

	sqlStmt := fmt.Sprintf("SELECT * FROM %s", TableUser)
	results, err := usr.App.DB.Query(sqlStmt)
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

func (usr MethodRepo) Delete(id int) (*model.User, error) {
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", TableUser)

	_, err := usr.App.DB.Exec(sqlStmt, id)
	CheckErr(err)

	return nil, err
}

func (usr MethodRepo) Update(id int, u *model.User) *model.User {

	sqlStmt := fmt.Sprintf("UPDATE %s SET name=?, email=?, phone_number=?, password=?, status=? WHERE id=%d ", TableUser, id)
	stmt, err := usr.App.DB.Prepare(sqlStmt)
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
*/