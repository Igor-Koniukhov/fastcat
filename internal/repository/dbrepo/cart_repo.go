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

type CartRepository interface {
	Create(cart *models.Cart) (*models.Cart, error)
	Get(id int) *models.Cart
	GetAll() []models.Cart
	Delete(id int) error
	Update(id int, u *models.Cart) *models.Cart
}

type CartRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewCartRepository(app *config.AppConfig, DB *sql.DB) *CartRepo {
	return &CartRepo{App: app, DB: DB}
}

func (c CartRepo) Create(cart *models.Cart) (*models.Cart, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("INSERT INTO %s (user_id, product_id, item) VALUES (?, ?, ?)", models.TableCarts)
		_, err := c.DB.ExecContext(ctx, sqlStmt, cart.UserID, cart.ProductID, cart.Items)
	web.Log.Error(err, err)
	return cart, err
}

func (c CartRepo) Get(id int) *models.Cart {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var cart models.Cart
	sqlStmt := fmt.Sprintf("SELECT id, user_id, product_id, item FROM %s WHERE id = ? ", models.TableCarts)
	err := c.DB.QueryRowContext(ctx, sqlStmt, id).Scan(
		&cart.ID,
		&cart.ProductID,
		&cart.Items)
	web.Log.Error(err, err)
	return &cart
}

func (c CartRepo) GetAll() []models.Cart {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var cart models.Cart
	var carts []models.Cart
	sqlStmt := fmt.Sprintf("SELECT id, user_id, product_id, items FROM %s", models.TableCarts)
	results, err := c.DB.QueryContext(ctx, sqlStmt)
	web.Log.Error(err, err)
	for results.Next() {
		err = results.Scan(
			&cart.ID,
			&cart.UserID,
			&cart.ProductID,
			&cart.Items)
		web.Log.Error(err, err)
		carts = append(carts, cart)
	}
	return carts
}

func (c CartRepo) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", models.TableCarts)
	_, err := c.DB.ExecContext(ctx, sqlStmt, id)
	web.Log.Error(err, err)
	return err
}

func (c CartRepo) Update(id int, cart *models.Cart) *models.Cart {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, user_id=?, cart_id=?, address_id=?, status=? WHERE id=%d ", models.TableCarts, id)
	_, err := c.DB.ExecContext(ctx, sqlStmt,
		cart.Items,
		cart.UserID,
		cart.ProductID,
		cart.Items)
	web.Log.Error(err, err)
	fmt.Println(cart)
	return cart
}
