package repository

import (
	"database/sql"
	"fmt"
	dr "github.com/igor-koniukhov/fastcat/driver"
	"strconv"
	"strings"

	"github.com/igor-koniukhov/fastcat/internal/model"
	"net/http"
)

type OrderRepositoryI interface {
	CreateOrder(ord *model.Order, db *sql.DB) (*model.Order, error)
	GetOrder(nameParam, param *string, db *sql.DB) *model.Order
	GetAllOrders( db *sql.DB) *[]model.Order
	DeleteOrder(id int , db *sql.DB) error
	UpdateOrder(id int, u *model.Order, db *sql.DB) *model.Order
	Param(r *http.Request) (string, string, int)
}
var order model.Order
type OrderRepository struct {}

func (o OrderRepository) CreateOrder(ord *model.Order, db *sql.DB) (*model.Order, error) {

	sqlStmt := fmt.Sprintf("INSERT INTO %s (user_id, cart_id, address, status) VALUES (?, ?, ?, ?)", dr.TableOrders)
	stmt, err := db.Prepare(sqlStmt)
	CheckErr(err)
	_, err = stmt.Exec(ord.UserID, ord.CartID, ord.AddressID, ord.Status)
	CheckErr(err)

	return ord, err
}


func (o OrderRepository) GetOrder( param *string, db *sql.DB) *model.Order {
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id = '%s' ", dr.TableOrders, *param)
	err := db.QueryRow(sqlStmt).Scan(
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

func (o OrderRepository) GetAllOrders(db *sql.DB) *[]model.Order {
	var orders []model.Order
	sqlStmt := fmt.Sprintf("SELECT * FROM %s", dr.TableUser)
	results, err := db.Query(sqlStmt)
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

func (o OrderRepository) DeleteOrder(id int, db *sql.DB) error {
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", dr.TableOrders)
	_, err := db.Exec(sqlStmt, id)
	CheckErr(err)
	return err
}

func (o OrderRepository) UpdateOrder(id int, ord *model.Order, db *sql.DB) *model.Order {
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, user_id=?, cart_id=?, address=?, status=? WHERE id=%d ", dr.TableOrders, id)
	stmt, err := db.Prepare(sqlStmt)
	CheckErr(err)
	_, err = stmt.Exec(
		order.ID,
		order.UserID,
		order.CartID,
		order.AddressID,
		order.Status)
	CheckErr(err)
	fmt.Println(*ord)
	return ord
}

func (o OrderRepository) Param(r *http.Request) (string, string, int) {
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

