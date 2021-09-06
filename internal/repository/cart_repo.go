package repository

import (
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"net/http"
	"strconv"
	"strings"
)

type CartRepositoryI interface {
	Create(u *model.Cart) (*model.Cart, error)
	Get(nameParam, param *string) *model.Cart
	GetAll() *[]model.Cart
	Delete(id int) error
	Update(id int, u *model.Cart) *model.Cart
	Param(r *http.Request) (string, string, int)
}

var RepoC *CartRepository

type CartRepository struct {
	App *config.AppConfig
}

func NewCartRepository(app *config.AppConfig) *CartRepository {
	return &CartRepository{App: app}
}

func NewRepoC(r *CartRepository) {
	RepoC = r
}

func (c CartRepository) Create(u *model.Cart) (*model.Cart, error) {

	return nil, nil
}

func (c CartRepository) Get(nameParam, param *string) *model.Cart {
	return nil
}

func (c CartRepository) GetAll() *[]model.Cart {
	return nil
}

func (c CartRepository) Delete(id int) error {
	return nil
}

func (c CartRepository) Update(id int, u *model.Cart, ) *model.Cart {
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
