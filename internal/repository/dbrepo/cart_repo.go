package dbrepo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	web "github.com/igor-koniukhov/webLogger/v2"
	"log"
	"time"
)

type CartRepository interface {
	Create(cart *models.CartResponse) (*models.CartResponse, int, error)
	Get(id int) *models.Cart
	GetAll() []models.CartResponse
	Delete(id int) error
	Update(id int, u *models.CartResponse) *models.CartResponse
	GetAllByUserID(id int) ([]models.Cart, error)
}

type CartRepo struct {
	DB  *sql.DB
	App *config.AppConfig
}

func NewCartRepository(app *config.AppConfig, DB *sql.DB) *CartRepo {
	return &CartRepo{App: app, DB: DB}
}

func (c CartRepo) Create(cart *models.CartResponse) (*models.CartResponse, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("INSERT INTO %s (user_id, user, address_delivery, cart_body, amount) VALUES (?, ?, ?, ?, ?)", models.TableCarts)
	result, err := c.DB.ExecContext(ctx, sqlStmt,
		&cart.UserID,
		&cart.User,
		&cart.AddressDelivery,
		&cart.CartBody,
		&cart.Amount)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	id, _ := result.LastInsertId()
	return cart, int(id), nil
}

func (c CartRepo) Get(id int) *models.Cart {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var cart models.CartResponse
	sqlStmt := fmt.Sprintf("SELECT id, user, address_delivery, cart_body, amount FROM %s WHERE id = ? ", models.TableCarts)
	err := c.DB.QueryRowContext(ctx, sqlStmt, id).Scan(
		&cart.ID,
		&cart.User,
		&cart.AddressDelivery,
		&cart.CartBody,
		&cart.Amount)
	if err != nil {
		log.Println(err)
	}
	var user *models.User
	var carts []models.CartBody
	err = json.Unmarshal(cart.User, &user)
	if err != nil {
		web.Log.Error(err, cart.User, user)
	}
	err = json.Unmarshal(cart.CartBody, &carts)
	if err != nil {
		web.Log.Error(err)
	}

	crt := &models.Cart{
		ID:              cart.ID,
		User:            *user,
		AddressDelivery: cart.AddressDelivery,
		CartBodies:      carts,
		Amount:          cart.Amount,
	}

	return crt
}

func (c CartRepo) GetAllByUserID(id int) ([]models.Cart, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var cart models.CartResponse
	var carts []models.Cart
	var user *models.User
	var cartBody []models.CartBody

	sqlStmt := fmt.Sprintf("SELECT id, user, address_delivery, cart_body, amount, status, created_at FROM %s WHERE user_id = ? ", models.TableCarts)
	result, err := c.DB.QueryContext(ctx, sqlStmt, id)
	if err != nil {
		web.Log.Info(err)
	}
	for result.Next() {
		err := result.Scan(
			&cart.ID,
			&cart.User,
			&cart.AddressDelivery,
			&cart.CartBody,
			&cart.Amount,
			&cart.Status,
			&cart.CreatedAt,
		)
		if err != nil {
			web.Log.Error(err)
			return nil, err
		}

		err = json.Unmarshal(cart.User, &user)
		if err != nil {
			web.Log.Error(err, cart.User, user)
		}
		err = json.Unmarshal(cart.CartBody, &cartBody)
		if err != nil {
			web.Log.Error(err)
		}

		crt := &models.Cart{
			ID:              cart.ID,
			User:            *user,
			AddressDelivery: cart.AddressDelivery,
			CartBodies:      cartBody,
			Amount:          cart.Amount,
			Status:          cart.Status,
			CreatedAt:       cart.CreatedAt,
		}
		carts = append(carts, *crt)
	}
	return carts, nil
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
			&cart.User,
			&cart.AddressDelivery,
			&cart.CartBody,
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
		cart.User,
		cart.AddressDelivery,
		cart.CartBody)
	if err != nil {
		log.Println(err)
	}
	return cart
}
