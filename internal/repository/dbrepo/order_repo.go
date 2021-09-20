package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"log"
	"time"
)

type OrderRepository interface {
	Create(or *models.Order) (*models.Order, error)
	Get(id int) *models.Order
	GetAll() *[]models.Order
	Delete(id int) error
	Update(id int, ord *models.Order) *models.Order
}

type OrderRepo struct {
	DB *sql.DB
	App *config.AppConfig
}

func NewOrderRepository(app *config.AppConfig, DB *sql.DB) *OrderRepo {
	return &OrderRepo{App:app, DB: DB}
}
func (o OrderRepo) Create(or *models.Order) (*models.Order, error) {
	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//defer cancel()
	/*sqlStmt := fmt.Sprintf("INSERT INTO %s (user_id, cart_id, address_id, status) VALUES (?, ?, ?, ?)", models.TableOrders)
	_, err := o.DB.ExecContext(ctx, sqlStmt, or.UserID, or.CartID, or.AddressID, or.Status)
	if err != nil {
		log.Println(err)
		return nil, err
	}*/
	fmt.Println(or)

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


		&order.Status,
		&order.CreatedAt,
		&order.UpdatedAt)
	if err != nil {
		log.Println(err)
	}
	return &order
}

func (o OrderRepo) GetAll() *[]models.Order {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var order models.Order
	var orders []models.Order
	sqlStmt := fmt.Sprintf("SELECT * FROM %s", models.TableOrders)
	results, err := o.DB.QueryContext(ctx, sqlStmt)
	if err != nil {
		log.Println(err)
	}
	for results.Next() {
		err = results.Scan(
			&order.ID,
			&order.UserID,


			&order.Status,
			&order.CreatedAt,
			&order.UpdatedAt)
		if err != nil {
			log.Println(err)
		}
		orders = append(orders, order)
	}
	return &orders
}

func (o OrderRepo) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", models.TableOrders)
	_, err := o.DB.ExecContext(ctx, sqlStmt, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (o OrderRepo) Update(id int, ord *models.Order) *models.Order {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, user_id=?, cart_id=?, address_id=?, status=? WHERE id=%d ", models.TableOrders, id)
	_, err := o.DB.ExecContext(ctx, sqlStmt,
		ord.ID,
		ord.UserID,


		ord.Status)
	if err != nil {
		log.Println(err)
	}
	return ord
}
