package repository

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"net/http"
)

type CartRepositoryI interface {
	CreateCart(u *model.Cart, db *sql.DB) (*model.Cart, error)
	GetCart(nameParam, param *string, db *sql.DB) *model.Cart
	GetAllUsers(db *sql.DB) *[]model.Cart
	DeleteUser(id int, db *sql.DB) error
	UpdateUser(id int, u *model.Cart, db *sql.DB) *model.Cart
	Param(r *http.Request) (string, string, int)
}

type CartRepository struct {
	App *config.AppConfig
}

func (c CartRepository) CreateCart(u *model.Cart, db *sql.DB) (*model.Cart, error) {

	return nil, nil
}

func (c CartRepository) GetCart(nameParam, param *string, db *sql.DB) *model.Cart {
	return nil
}

func (c CartRepository) GetAllUsers(db *sql.DB) *[]model.Cart {
	return nil
}

func (c CartRepository) DeleteUser(id int, db *sql.DB) error {
	return nil
}

func (c CartRepository) UpdateUser(id int, u *model.Cart, db *sql.DB) *model.Cart {
	return nil
}

func (c CartRepository) Param(r *http.Request) (string, string, int) {
	return "", "", 0
}
