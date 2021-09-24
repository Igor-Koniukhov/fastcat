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

type ProductRepository interface {
	Create(item *models.Item, id int) (*models.Item, error)
	GetBySchedule(id int) []models.Item
	Get(id int) *models.Item
	GetAll() []models.Item
	GetAllBySupplierID(id int) []models.Item
	Delete(id int) (err error)
	SoftDelete(id int) error
	Update(id int, item *models.Item, ) *models.Item
}

type ProductRepo struct {
	DB  *sql.DB
	App *config.AppConfig
}

func NewProductRepository(app *config.AppConfig, DB *sql.DB) *ProductRepo {
	return &ProductRepo{App: app, DB: DB}
}

func (p ProductRepo) Create(item *models.Item, id int) (*models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmtSQl := fmt.Sprintf("INSERT INTO %s (image, ingredients, name, price, type,  supplier_id) VALUES (?, ?, ?, ?, ?, ?)", models.TabItems)
	ingredients, err := json.MarshalIndent(item.Ingredients, "", "")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	stmtRest := fmt.Sprintf("INSERT %s SET  item_id=? ", models.TabSuppliersItems)
	result, err := p.DB.ExecContext(ctx, stmtSQl,
		&item.Image,
		ingredients,
		&item.Name,
		&item.Price,
		&item.Type,
		id,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_, err = p.DB.ExecContext(ctx, stmtRest, int(lastInsertedID))
	web.Log.Info(id, "- product id ", lastInsertedID, " -item id ")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return item, nil
}

func (p ProductRepo) Get(id int) *models.Item {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var product models.Product
	var item models.Item
	sqlStmt := fmt.Sprintf("SELECT id, image, ingredients, name, price, type FROM %s WHERE id = ? ", models.TabItems)
	err := p.DB.QueryRowContext(ctx, sqlStmt, id).Scan(
		&product.Id,
		&product.Image,
		&product.Ingredients,
		&product.Name,
		&product.Price,
		&product.Type,
	)
	if err != nil {
		log.Println(err)
	}
	str := []string{string(product.Ingredients)}

	err = json.Unmarshal(product.Ingredients, &str)
	if err != nil {
		log.Println(err)
	}
	item = models.Item{
		Id:          product.Id,
		Image:       product.Image,
		Ingredients: str,
		Name:        product.Name,
		Price:       product.Price,
		Type:        product.Type,
	}
	if err != nil {
		log.Println(err)
	}
	return &item
}

func (p ProductRepo) GetBySchedule(id int) []models.Item {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var product models.Product
	var items []models.Item
	var str []string
	sqlStmt := fmt.Sprintf("SELECT id, image, ingredients, name, price, type, supplier_id FROM %s WHERE deleted_at IS NULL ", models.TabItems)

	results, err := p.DB.QueryContext(ctx, sqlStmt)
	if err != nil {
		log.Println(err)
	}
	for results.Next() {
		err = results.Scan(
			&product.Id,
			&product.Image,
			&product.Ingredients,
			&product.Name,
			&product.Price,
			&product.Type,
			&product.SuppliersID,
		)
		if err != nil {
			log.Println(err)
		}
		err := json.Unmarshal(product.Ingredients, &str)
		if err != nil {
			log.Println(err)
		}
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

func (p ProductRepo) GetAll() []models.Item {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var product models.Product
	var items []models.Item
	var str []string
	sqlStmt := fmt.Sprintf("SELECT id, image, ingredients, name, price, type, supplier_id FROM %s WHERE deleted_at IS NULL", models.TabItems)

	results, err := p.DB.QueryContext(ctx, sqlStmt)
	if err != nil {
		log.Println(err)
	}
	for results.Next() {
		err = results.Scan(
			&product.Id,
			&product.Image,
			&product.Ingredients,
			&product.Name,
			&product.Price,
			&product.Type,
			&product.SuppliersID,
		)
		if err != nil {
			log.Println(err)
		}
		err := json.Unmarshal(product.Ingredients, &str)
		if err != nil {
			log.Println(err)
		}
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
func (p ProductRepo) GetAllBySupplierID(id int) []models.Item {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var product models.Product
	var items []models.Item
	var str []string
	sqlStmt := fmt.Sprintf("SELECT id, image, ingredients, name, price, type, supplier_id FROM %s WHERE deleted_at IS NULL AND supplier_id=?", models.TabItems)

	results, err := p.DB.QueryContext(ctx, sqlStmt, id)
	if err != nil {
		log.Println(err)
	}
	for results.Next() {
		err = results.Scan(
			&product.Id,
			&product.Image,
			&product.Ingredients,
			&product.Name,
			&product.Price,
			&product.Type,
			&product.SuppliersID,
		)
		if err != nil {
			log.Println(err)
		}
		err := json.Unmarshal(product.Ingredients, &str)
		if err != nil {
			log.Println(err)
		}
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=? ", models.TabItems)
	_, err = p.DB.ExecContext(ctx, sqlStmt, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (p ProductRepo) SoftDelete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	TimeFormat := time.Now().UTC().Format("2006-01-02 15:04:05.999999")
	sqlStmt := fmt.Sprintf("UPDATE %s SET deleted_at = ? WHERE supplier_id = ?", models.TabItems)
	_, err := p.DB.ExecContext(ctx, sqlStmt, TimeFormat, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (p ProductRepo) Update(id int, item *models.Item) *models.Item {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, name=?, price=?, image=?, type=?, ingredienst=? , WHERE id=?", models.TabItems)
	_, err := p.DB.ExecContext(ctx, sqlStmt, item.Id, item.Name, item.Price, item.Image, item.Type, item.Ingredients, id)
	if err != nil {
		log.Println(err)
	}
	return item
}
