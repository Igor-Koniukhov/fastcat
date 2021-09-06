package repository

import (
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"net/http"
	"strconv"
	"strings"
)

type CartRepositoryI interface {
	CreateCart(u *model.Cart) (*model.Cart, error)
	GetCart(nameParam, param *string) *model.Cart
	GetAllCarts() *[]model.Cart
	DeleteCarts(id int) error
	UpdateCarts(id int, u *model.Cart) *model.Cart
	Param(r *http.Request) (string, string, int)
}

var RepoC *CartRepository

type CartRepository struct {
	App *config.AppConfig
}

func NewCartRepository(app *config.AppConfig) *CartRepository {
	return &CartRepository{App: app}
}

func NewRepoC(r *CartRepository)  {

	RepoC = r

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

func (c CartRepository) DeleteCarts(id int) error {
	return nil
}

func (c CartRepository) UpdateCarts(id int, u *model.Cart, ) *model.Cart {
	return nil
}

func (c CartRepository) Param(r *http.Request) (string, string, int) {
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
