package repository

import (
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	web "github.com/igor-koniukhov/webLogger/v3"
)

type CartRepositoryInterface interface {
	Create(cart *model.Cart) (*model.Cart, error)
	Get(nameParam, param *string) *model.Cart	
	GetAll() *[]model.Cart
	Delete(id int) error
	Update(id int, u *model.Cart) *model.Cart
}

var cart model.Cart
type CartRepository struct{
	App *config.AppConfig
}

func NewCartRepository(app *config.AppConfig) *CartRepository {
	return &CartRepository{App: app}
}

func (c CartRepository) Create(cart *model.Cart) (*model.Cart, error) {
	sqlStmt := fmt.Sprintf("INSERT INTO %s (user_id, product_id, item) VALUES (?, ?, ?)", model.TableCarts)
	p, err := c.App.DB.Prepare(sqlStmt)
	defer p.Close()
	web.Log.Error(err, err)
	_, err = p.Exec( cart.UserID, cart.ProductID, cart.Items )
	web.Log.Error(err, err)
	return cart, err
}

func (c CartRepository) Get(nameParam, param *string) *model.Cart {
	sqlStmt := fmt.Sprintf("SELECT id, user_id, product_id, item FROM %s WHERE id = '%s' ", model.TableCarts, *param)
	err := c.App.DB.QueryRow(sqlStmt).Scan(
		&cart.ID,
		&cart.ProductID,
		&cart.Items)
	web.Log.Error(err, err)
	return &cart
}

func (c CartRepository) GetAll() *[]model.Cart {
	var carts []model.Cart
	sqlStmt := fmt.Sprintf("SELECT id, user_id, product_id, items FROM %s", model.TableCarts)
	results, err := c.App.DB.Query(sqlStmt)
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
	return &carts
}

func (c CartRepository) Delete(id int) error {
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", model.TableCarts)
	_, err := c.App.DB.Exec(sqlStmt, id)
	web.Log.Error(err, err)
	return err
}

func (c CartRepository) Update(id int, cart *model.Cart) *model.Cart {
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, user_id=?, cart_id=?, address_id=?, status=? WHERE id=%d ", model.TableCarts, id)
	fmt.Println(sqlStmt)
	stmt, err := c.App.DB.Prepare(sqlStmt)
	web.Log.Error(err, err)
	_, err = stmt.Exec(
		cart.Items,
		cart.UserID,
		cart.ProductID,
		cart.Items)
	web.Log.Error(err, err)
	fmt.Println(cart)
	return cart
}

