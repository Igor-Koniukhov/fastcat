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
	Create(ord *models.Order) (*models.Order, error)
	Get(id int) *models.Order
	GetAll() *[]models.Order
	GetAllByUserID(id int) []models.Cart
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
func (o OrderRepo) Create(ord *models.Order) (*models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("INSERT INTO %s (supplier_id, cart_id, title, price, quantity, status, ) VALUES (?, ?, ?, ?, ?, ?)", models.TableOrders)
	_, err := o.DB.ExecContext(ctx, sqlStmt,
		ord.SupplierId,
		ord.BodyOrder.CartId,
		ord.BodyOrder.Title,
		ord.BodyOrder.Price,
		ord.BodyOrder.Quantity,
		ord.BodyOrder.Status)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return ord, nil
}

func (o OrderRepo) Get(id int) *models.Order {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var ord models.Order
	sqlStmt := fmt.Sprintf("SELECT id, supplier_id, cart_id, title, price, quantity, status FROM %s WHERE supplier_id = ? ", models.TableOrders)
	err := o.DB.QueryRowContext(ctx, sqlStmt, id).Scan(
		&ord.ID,
		&ord.SupplierId,
		&ord.BodyOrder.CartId,
		&ord.BodyOrder.Title,
		&ord.BodyOrder.Price,
		&ord.BodyOrder.Quantity,
		&ord.BodyOrder.Status)
	if err != nil {
		log.Println(err)
	}
	return &ord
}

func (o OrderRepo) GetAll() *[]models.Order {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var ord models.Order
	var orders []models.Order
	sqlStmt := fmt.Sprintf("SELECT id, supplier_id, cart_id, title, price, quantity, status FROM %s", models.TableOrders)
	results, err := o.DB.QueryContext(ctx, sqlStmt)
	if err != nil {
		log.Println(err)
	}
	for results.Next() {
		err = results.Scan(
			&ord.ID,
			&ord.SupplierId,
			&ord.BodyOrder.CartId,
			&ord.BodyOrder.Title,
			&ord.BodyOrder.Price,
			&ord.BodyOrder.Quantity,
			&ord.BodyOrder.Status)
		if err != nil {
			log.Println(err)
		}
		orders = append(orders, ord)
	}
	return &orders
}
func (o OrderRepo) GetAllByUserID(id int) []models.Cart {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var cart models.Cart
	var carts []models.Cart
	sqlStmt := fmt.Sprintf("SELECT id, user, address_delivery, cart_body, amount, status, created_at, updated_at status FROM %s WHERE user_id=? AND status = 'accepted' ", models.TableCarts)
	results, err := o.DB.QueryContext(ctx, sqlStmt)
	if err != nil {
		log.Println(err)
	}
	for results.Next() {
		err = results.Scan(
			&cart.ID,
			&cart.User,
			&cart.UserID,
			&cart.AddressDelivery,
			&cart.CartBodies,
			&cart.Amount,
			&cart.Status,
			&cart.CreatedAt,
			)
		if err != nil {
			log.Println(err)
		}
		carts = append(carts, cart)
	}
	return carts
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
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, supplier_id=?, cart_id=?, cart_id=?, title=? price=?, quantity=?, status=? WHERE id=%d ", models.TableOrders, id)
	_, err := o.DB.ExecContext(ctx, sqlStmt,
		ord.ID,
		ord.SupplierId,
		ord.BodyOrder.CartId,
		ord.BodyOrder.Title,
		ord.BodyOrder.Price,
		ord.BodyOrder.Quantity,
		ord.BodyOrder.Status)
	if err != nil {
		log.Println(err)
	}
	return ord
}
