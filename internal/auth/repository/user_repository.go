package repository

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gopath/auth/model"
)

type IUserRepository interface {
	GetUserByEmail(email string) (*model.User, error)
	GetUserByID(id int) (*model.User, error)
}
type UserRepository struct {
	users []*model.User
}

func NewUserRepository() IUserRepository {
	p1, _ := bcrypt.GenerateFromPassword([]byte("11111111"), bcrypt.DefaultCost)
	p2, _ := bcrypt.GenerateFromPassword([]byte("22222222"), bcrypt.DefaultCost)

	users := []*model.User{
		&model.User{
			ID:       116,
			Email:    "ikoniukhov@gmail.com",
			Name:     "",
			Password: string(p1),
		},
		&model.User{
			ID:       2,
			Email:    "mary@gmail.com",
			Name:     "Mary",
			Password: string(p2),
		},
	}

	return &UserRepository{users: users}

}

func (r UserRepository) GetUserByEmail(email string) (*model.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r UserRepository) GetUserByID(id int) (*model.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
