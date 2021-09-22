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

type CartRepository interface {
	Create(cart *models.CartResponse) (*models.CartResponse, error)
	Get(id int) *models.CartResponse
	GetAll() []models.CartResponse
	Delete(id int) error
	Update(id int, u *models.CartResponse) *models.CartResponse
}

type CartRepo struct {
	DB  *sql.DB
	App *config.AppConfig
}

func NewCartRepository(app *config.AppConfig, DB *sql.DB) *CartRepo {
	return &CartRepo{App: app, DB: DB}
}

func (c CartRepo) Create(cart *models.CartResponse) (*models.CartResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("INSERT INTO %s (user_id, address_delivery, order_body, amount) VALUES (?, ?, ?, ?)", models.TableCarts)
		_, err := c.DB.ExecContext(ctx, sqlStmt,
			&cart.User.ID,
			&cart.AddressDelivery,
			&cart.OrderBody,
			&cart.Amount)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return cart, nil
}

func (c CartRepo) Get(id int) *models.CartResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var cart models.CartResponse
	sqlStmt := fmt.Sprintf("SELECT id, user_id, address_delivery, order_body, amount FROM %s WHERE id = ? ", models.TableCarts)
	err := c.DB.QueryRowContext(ctx, sqlStmt, id).Scan(
		&cart.ID,
		&cart.User.ID,
		&cart.AddressDelivery,
		&cart.OrderBody,
		&cart.Amount)
	if err != nil {
		log.Println(err)
	}
	return &cart
}

func (c CartRepo) GetAll() []models.CartResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var cart models.CartResponse
	var carts []models.CartResponse
	sqlStmt := fmt.Sprintf("SELECT id, user_id, address_delivery, order_body, amount FROM %s", models.TableCarts)
	results, err := c.DB.QueryContext(ctx, sqlStmt)
	if err != nil {
		log.Println(err)
	}
	for results.Next() {
		err = results.Scan(
			&cart.ID,
			&cart.User.ID,
			&cart.AddressDelivery,
			&cart.OrderBody,
			&cart.Amount)
		if err != nil {
			log.Println(err)
		}
		carts = append(carts, cart)
	}
	return carts
}

func (c CartRepo) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", models.TableCarts)
	_, err := c.DB.ExecContext(ctx, sqlStmt, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c CartRepo) Update(id int, cart *models.CartResponse) *models.CartResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, user_id=?, address_delivery=?, order_body=? WHERE id=%d ", models.TableCarts, id)
	_, err := c.DB.ExecContext(ctx, sqlStmt,
		cart.ID,
		cart.User.ID,
		cart.AddressDelivery,
		cart.OrderBody)
	if err != nil {
		log.Println(err)
	}
	return cart
}
