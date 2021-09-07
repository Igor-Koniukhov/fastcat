package repository

import (
	"fmt"
	dr "github.com/igor-koniukhov/fastcat/driver"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	web "github.com/igor-koniukhov/webLogger/v3"
)


type OrderRepositoryInterface interface {
	Create(or *model.Order) (*model.Order, error)
	Get(id int) *model.Order
	GetAll() *[]model.Order
	Delete(id int) error
	Update(id int, ord *model.Order) *model.Order
}
var order model.Order

type OrderRepository struct{

	App *config.AppConfig
}

func NewOrderRepository(app *config.AppConfig) *OrderRepository {
	return &OrderRepository{App: app}
}
func (o OrderRepository) Create(or *model.Order) (*model.Order, error) {
	sqlStmt := fmt.Sprintf("INSERT INTO %s (user_id, cart_id, address_id, status) VALUES (?, ?, ?, ?)", dr.TableOrders)
	p, err := o.App.DB.Prepare(sqlStmt)
	defer p.Close()
	web.Log.Error(err, err)
	_, err = p.Exec(or.UserID, or.CartID, or.AddressID, or.Status)
	web.Log.Error(err, err)
	return or, err
}

func (o OrderRepository) Get(id int) *model.Order {
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id = ? ", dr.TableOrders)
	err := o.App.DB.QueryRow(sqlStmt, id).Scan(
		&order.ID,
		&order.UserID,
		&order.CartID,
		&order.AddressID,
		&order.Status,
		&order.CreatedAt,
		&order.UpdatedAt)
	web.Log.Error(err, err)
	return &order
}

func (o OrderRepository) GetAll() *[]model.Order {
	var orders []model.Order
	sqlStmt := fmt.Sprintf("SELECT * FROM %s", dr.TableOrders)
	results, err := o.App.DB.Query(sqlStmt)
	web.Log.Error(err, err)
	for results.Next() {
		err = results.Scan(
			&order.ID,
			&order.UserID,
			&order.CartID,
			&order.AddressID,
			&order.Status,
			&order.CreatedAt,
			&order.UpdatedAt)
		web.Log.Error(err, err)
		orders = append(orders, order)
	}
	return &orders
}

func (o OrderRepository) Delete(id int) error {
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", dr.TableOrders)
	_, err := o.App.DB.Exec(sqlStmt, id)
	web.Log.Error(err, err)
	return err
}

func (o OrderRepository) Update(id int, ord *model.Order) *model.Order {
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, user_id=?, cart_id=?, address_id=?, status=? WHERE id=%d ", dr.TableOrders, id)
	fmt.Println(sqlStmt)
	stmt, err := o.App.DB.Prepare(sqlStmt)
	web.Log.Error(err, err)
	_, err = stmt.Exec(
		ord.ID,
		ord.UserID,
		ord.CartID,
		ord.AddressID,
		ord.Status)
	web.Log.Error(err, err)
	fmt.Println(*ord)
	return ord
}






