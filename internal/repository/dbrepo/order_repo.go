package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	web "github.com/igor-koniukhov/webLogger/v3"
	"time"
)

type OrderRepository interface {
	Create(or *models.Order) (*models.Order, error)
	Get(id int) *models.Order
	GetAll() *[]models.Order
	Delete(id int) error
	Update(id int, ord *models.Order) *models.Order
}

type OrderRepo struct{
	App *config.AppConfig
	DB *sql.DB
}

func NewOrderRepository(app *config.AppConfig, DB *sql.DB) *OrderRepo {
	return &OrderRepo{App: app, DB: DB}
}
func (o OrderRepo) Create(or *models.Order) (*models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("INSERT INTO %s (user_id, cart_id, address_id, status) VALUES (?, ?, ?, ?)", models.TableOrders)
		_, err := o.DB.ExecContext(ctx, sqlStmt, or.UserID, or.CartID, or.AddressID, or.Status)
	web.Log.Error(err, err)
	return or, nil
}

func (o OrderRepo) Get(id int) *models.Order {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var order models.Order
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE id = ? ", models.TableOrders)
	err := o.DB.QueryRowContext(ctx, sqlStmt, id).Scan(
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

func (o OrderRepo) GetAll() *[]models.Order {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var order models.Order
	var orders []models.Order
	sqlStmt := fmt.Sprintf("SELECT * FROM %s", models.TableOrders)
	results, err := o.DB.QueryContext(ctx, sqlStmt)
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

func (o OrderRepo) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", models.TableOrders)
	_, err := o.DB.ExecContext(ctx, sqlStmt, id)
	web.Log.Error(err, err)
	return err
}

func (o OrderRepo) Update(id int, ord *models.Order) *models.Order {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, user_id=?, cart_id=?, address_id=?, status=? WHERE id=%d ", models.TableOrders, id)
	_, err := o.DB.ExecContext(ctx, sqlStmt,
		ord.ID,
		ord.UserID,
		ord.CartID,
		ord.AddressID,
		ord.Status)
	web.Log.Error(err, err)
	fmt.Println(*ord)
	return ord
}






