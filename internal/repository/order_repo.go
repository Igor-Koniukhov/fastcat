package repository

import (
	"github.com/igor-koniukhov/fastcat/internal/model"
	"net/http"
)

type OrderRepositoryI interface {
	CreateOrder(u *model.Order) (*model.Order, error)
	GetOrder(nameParam, param *string) *model.Order
	GetAllOrders() *[]model.Order
	DeleteOrder(id int ) error
	UpdateOrder(id int, u *model.Order) *model.Order
	Param(r *http.Request) (string, string, int)
}

type OrderRepository struct {}

func (o OrderRepository) CreateOrder(u *model.Order) (*model.Order, error) {
	return nil, nil
}


func (o OrderRepository) GetOrder(nameParam, param *string) *model.Order {
	return nil
}

func (o OrderRepository) GetAllOrders() *[]model.Order {
	return nil
}

func (o OrderRepository) DeleteOrder(id int) error {
	return nil
}

func (o OrderRepository) UpdateOrder(id int, u *model.Order) *model.Order {
	return nil
}

func (o OrderRepository) Param(r *http.Request) (string, string, int) {
	return "", "", 0
}

