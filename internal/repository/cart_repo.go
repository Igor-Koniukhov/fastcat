package repository

import (
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"net/http"
)

type CartRepositoryI interface {
	CreateCart(u *model.Cart) (*model.Cart, error)
	GetCart(nameParam, param *string) *model.Cart
	GetAllCarts() *[]model.Cart
	DeleteCarts(id int) error
	UpdateCarts(id int, u *model.Cart) *model.Cart
	Param(r *http.Request) (string, string, int)
}

type CartRepository struct {
	App *config.AppConfig
}

func (c CartRepository) CreateCart(u *model.Cart) (*model.Cart, error) {

	return nil, nil
}

func (c CartRepository) GetCart(nameParam, param *string) *model.Cart {
	return nil
}

func (c CartRepository) GetAllCarts() *[]model.Cart {
	return nil
}

func (c CartRepository) DeleteCarts(id int ) error {
	return nil
}

func (c CartRepository) UpdateCarts(id int, u *model.Cart, ) *model.Cart {
	return nil
}

func (c CartRepository) Param(r *http.Request) (string, string, int) {
	return "", "", 0
}
