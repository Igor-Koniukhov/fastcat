package dbrepo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	web "github.com/igor-koniukhov/webLogger/v3"
	"time"
)


type ProductRepository interface {
	Create(item *models.Item, id int) (*models.Item, error)
	Get(id int) *models.Item
	GetAll() []models.Item
	Delete(id int) (err error)
	SoftDelete(id int) error
	Update(id int, item *models.Item, ) *models.Item
}

type ProductRepo struct{
	App *config.AppConfig
	DB *sql.DB
}

func NewProductRepository(app *config.AppConfig, DB *sql.DB) *ProductRepo {
	return &ProductRepo{App: app, DB: DB}
}

func (p ProductRepo) Create(item *models.Item, id int) (*models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmtSQl := fmt.Sprintf("INSERT INTO %s (name, price, type, image, ingredients, supplier_id) VALUES (?, ?, ?, ?, ?, ?)", models.TabItems)
	ingredients, err := json.MarshalIndent(item.Ingredients, "", "")

	stmtRest := fmt.Sprintf("INSERT %s SET  item_id=? ", models.TabSuppliersItems)
	result, err := p.DB.ExecContext(ctx, stmtSQl,
		&item.Name,
		&item.Price,
		&item.Type,
		&item.Image,
		ingredients,
		id,
	)
	web.Log.Error(err)
	lastInsertedID, err := result.LastInsertId()
	web.Log.Error(err)

	_, err = p.DB.ExecContext(ctx, stmtRest, int(lastInsertedID))

	web.Log.Info(id, "- product id ", lastInsertedID," -item id " )
	web.Log.Error(err, err)

	return item, err
}

func (p ProductRepo) Get(id int) *models.Item {
	var product models.Product
	var item models.Item
	sqlStmt := fmt.Sprintf("SELECT id, name, price, image, type, ingredients FROM %s WHERE id = ? ", models.TabItems)
	err := p.DB.QueryRow(sqlStmt, id).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Image,
		&product.Type,
		&product.Ingredients,
	)
	str := []string{string(product.Ingredients)}

	json.Unmarshal(product.Ingredients, &str)
	item = models.Item{
		Id:          product.Id,
		Name:        product.Name,
		Price:       product.Price,
		Image:       product.Image,
		Type:        product.Type,
		Ingredients: str,
	}
	web.Log.Error(err, err)
	return &item
}

func (p ProductRepo) GetAll() []models.Item {
	var product models.Product
	var items []models.Item
	var str []string
	sqlStmt := fmt.Sprintf("SELECT id, name, price, image, type, ingredients, supplier_id FROM %s WHERE deleted_at IS NULL", models.TabItems)

	results, err := p.DB.Query(sqlStmt)
	web.Log.Error(err, err)
	for results.Next() {
		err = results.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Image,
			&product.Type,
			&product.Ingredients,
			&product.SuppliersID,
		)
		web.Log.Error(err, err)
		json.Unmarshal(product.Ingredients, &str)
		items = append(items, models.Item{
			Id:          product.Id,
			Name:        product.Name,
			Price:       product.Price,
			Image:       product.Image,
			Type:        product.Type,
			Ingredients: str,
			SuppliersID: product.SuppliersID,
		})
	}
	return items
}

func (p ProductRepo) Delete(id int) (err error) {
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=? ", models.TabItems)
	_, err = p.DB.Exec(sqlStmt, id)
	web.Log.Error(err, err)
	return nil
}

func (p ProductRepo) SoftDelete(id int) error {
	sqlStmt := fmt.Sprintf("UPDATE %s SET deleted_at = ? WHERE supplier_id = ?", models.TabItems)
	_, err := p.DB.Exec(sqlStmt, p.App.TimeFormat, id)
	web.Log.Error(err, err)
	return nil
}

func (p ProductRepo) Update(id int, item *models.Item) *models.Item {
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, name=?, price=?, image=?, type=?, ingredienst=? , WHERE id=?", models.TabItems)
	_, err := p.DB.Exec(sqlStmt, item.Id, item.Name, item.Price, item.Image, item.Type, item.Ingredients, id)
	web.Log.Error(err, err)
	return item
}




