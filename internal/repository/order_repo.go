package repository

import (

	"fmt"
	dr "github.com/igor-koniukhov/fastcat/driver"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"strconv"
	"strings"

	"github.com/igor-koniukhov/fastcat/internal/model"
	"net/http"
)

type OrderRepositoryI interface {
	CreateOrder(ord *model.Order) (*model.Order, error)
	GetOrder(nameParam, param *string) *model.Order
	GetAllOrders() *[]model.Order
	DeleteOrder(id int) error
	UpdateOrder(id int, u *model.Order) *model.Order
	Param(r *http.Request) (string, string, int)
}


var order model.Order
var RepoO *OrderRepository


type OrderRepository struct {
	App *config.AppConfig
}

func NewOrderRepository(app *config.AppConfig) *OrderRepository {
	return &OrderRepository{App: app}
}
func NewRepoO(r *OrderRepository) {
	RepoO = r

}

func (o *OrderRepository) CreateOrder(or *model.Order) (*model.Order, error) {

	sqlStmt := fmt.Sprintf("INSERT INTO %s (user_id, cart_id, address_id, status) VALUES (?, ?, ?, ?)", dr.TableOrders)

	p, err := o.App.DB.Prepare(sqlStmt)
	defer p.Close()
	CheckErr(err)
	_, err = p.Exec(or.UserID, or.CartID, or.AddressID, or.Status)
	CheckErr(err)
	return or, err
}

func (o *OrderRepository) GetOrder(param *string) *model.Order {
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id = '%s' ", dr.TableOrders, *param)
	err := o.App.DB.QueryRow(sqlStmt).Scan(
		&order.ID,
		&order.UserID,
		&order.CartID,
		&order.AddressID,
		&order.Status,
		&order.CreatedAt,
		&order.UpdatedAt)
	CheckErr(err)
	return &order
}

func (o *OrderRepository) GetAllOrders() *[]model.Order {
	var orders []model.Order
	sqlStmt := fmt.Sprintf("SELECT * FROM %s", dr.TableOrders)
	results, err := o.App.DB.Query(sqlStmt)
	CheckErr(err)
	for results.Next() {
		err = results.Scan(
			&order.ID,
			&order.UserID,
			&order.CartID,
			&order.AddressID,
			&order.Status,
			&order.CreatedAt,
			&order.UpdatedAt)
		CheckErr(err)
		orders = append(orders, order)
	}
	return &orders
}

func (o *OrderRepository) DeleteOrder(id int) error {
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", dr.TableOrders)
	_, err := o.App.DB.Exec(sqlStmt, id)
	CheckErr(err)
	return err
}

func (o *OrderRepository) UpdateOrder(id int, ord *model.Order) *model.Order {
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, user_id=?, cart_id=?, address_id=?, status=? WHERE id=%d ", dr.TableOrders, id)
	fmt.Println(sqlStmt)
	stmt, err := o.App.DB.Prepare(sqlStmt)
	CheckErr(err)
	_, err = stmt.Exec(
		ord.ID,
		ord.UserID,
		ord.CartID,
		ord.AddressID,
		ord.Status)
	CheckErr(err)
	fmt.Println(*ord)
	return ord
}

func (o *OrderRepository) Param(r *http.Request) (string, string, int) {
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
