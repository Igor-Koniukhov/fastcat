package repository

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"net/http"
)

type OrderRepositoryI interface {
	CreateOrder(u *model.Order, db *sql.DB) (*model.Order, error)
	GetOrder(nameParam, param *string, db *sql.DB) *model.Order
	GetAllOrders(db *sql.DB) *[]model.Order
	DeleteOrder(id int, db *sql.DB) error
	UpdateOrder(id int, u *model.Order, db *sql.DB) *model.Order
	Param(r *http.Request) (string, string, int)
}

type OrderRepository struct {
	App *config.AppConfig
}

func (o OrderRepository) CreateOrder(u *model.Order, db *sql.DB) (*model.Order, error) {
	return nil, nil
}


func (o OrderRepository) GetOrder(nameParam, param *string, db *sql.DB) *model.Order {
	return nil
}

func (o OrderRepository) GetAllOrders(db *sql.DB) *[]model.Order {
	return nil
}

func (o OrderRepository) DeleteOrder(id int, db *sql.DB) error {
	return nil
}

func (o OrderRepository) UpdateOrder(id int, u *model.Order, db *sql.DB) *model.Order {
	return nil
}

func (o OrderRepository) Param(r *http.Request) (string, string, int) {
	return "", "", 0
}

