package repository

import (
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	web "github.com/igor-koniukhov/webLogger/v3"
	"net/http"
	"strconv"
	"strings"
)

type ProductRepositoryI interface {
	Create(item *model.Item) (*model.Item, error)
	Get(id int) *model.Item
	GetAllP() *[]model.Item
	Delete(id int) error
	Update(id int, u *model.Product) *model.Item
	Param(r *http.Request) (string, string, int)
}

var RepoP *ProductRepository

type ProductRepository struct {
	App *config.AppConfig
}

func NewProductRepository(app *config.AppConfig) *ProductRepository {
	return &ProductRepository{App: app}
}

func NewRepoP(r *ProductRepository) {
	RepoP = r
}

func (p *ProductRepository) Create(item *model.Item, id int) (*model.Item, error) {
	stmtSQl := fmt.Sprintf("INSERT INTO %s (name, price, type, image, ingredients, supplier_id) VALUES (?, ?, ?, ?, ?, ?)", model.TabItems)
	ingredients, err := json.MarshalIndent(item.Ingredients, "", "")
	stmt, err := p.App.DB.Prepare(stmtSQl)
	defer stmt.Close()
	web.Log.Error(err)

	stmtRestaurantsItem := fmt.Sprintf("INSERT %s SET  item_id=? ", model.TabSuppliersItems)
	stmtRest, err := p.App.DB.Prepare(stmtRestaurantsItem)
	web.Log.Error(err, err)
	defer stmtRest.Close()

	result, err := stmt.Exec(
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

	_, err = stmtRest.Exec(int(lastInsertedID))

	web.Log.Info(id, "- product id ", lastInsertedID," -item id " )
	web.Log.Error(err, err)

	return item, err
}

func (p *ProductRepository) Get(id int) model.Item {
	var product model.Product
	var item model.Item
	sqlStmt := fmt.Sprintf("SELECT id, name, price, image, type, ingredients FROM %s WHERE id = %d ", model.TabItems, id)
	err := p.App.DB.QueryRow(sqlStmt).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Image,
		&product.Type,
		&product.Ingredients,
	)
	str := []string{string(product.Ingredients)}

	json.Unmarshal(product.Ingredients, &str)
	item = model.Item{
		Id:          product.Id,
		Name:        product.Name,
		Price:       product.Price,
		Image:       product.Image,
		Type:        product.Type,
		Ingredients: str,
	}
	CheckErr(err)
	return item
}

func (p *ProductRepository) GetAll() *[]model.Item {
	var product model.Product
	var items []model.Item
	var str []string
	sqlStmt := fmt.Sprintf("SELECT id, name, price, image, type, ingredients, supplier_id FROM %s", model.TabItems)

	results, err := p.App.DB.Query(sqlStmt)
	CheckErr(err)
	for results.Next() {
		err = results.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Image,
			&product.Type,
			&product.Ingredients,

		)

		CheckErr(err)
		json.Unmarshal(product.Ingredients, &str)
		items = append(items, model.Item{
			Id:          product.Id,
			Name:        product.Name,
			Price:       product.Price,
			Image:       product.Image,
			Type:        product.Type,
			Ingredients: str,
		})
	}

	return &items
}

func (p *ProductRepository) Delete(id int) (err error) {
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=? ", model.TabItems)
	_, err = p.App.DB.Exec(sqlStmt, id)
	fmt.Println(sqlStmt)
	web.Log.Error(err, err)
	return
}
func (p *ProductRepository) SoftDelete(id int) error {
	sqlStmt := fmt.Sprintf("UPDATE %s SET deleted_at = ? WHERE supplier_id = ?", model.TabItems)
	_, err := p.App.DB.Exec(sqlStmt, p.App.TimeFormat, id)
	web.Log.Error(err, err)
	return nil
}

func (p *ProductRepository) Update(id int, item *model.Item, ) *model.Item {
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, name=?, price=?, image=?, type=?, ingredienst=? , WHERE id=?", model.TabItems)
	_, err := p.App.DB.Exec(sqlStmt, item.Id, item.Name, item.Price, item.Image, item.Type, item.Ingredients, id)
	web.Log.Error(err, err)
	return item
}

func (p *ProductRepository) Param(r *http.Request) (string, string, int) {
	var paramName string
	var param string
	var id int
	fields := strings.Split(r.URL.String(), "/")
	str := fields[len(fields)-1]
	//TODO: need to be rewriting with regexp
	if len(str) > 5 {
		paramName = "email"
		param = str
		id = 0
	} else {
		num, err := strconv.Atoi(str)
		CheckErr(err)
		paramName = "id"
		param = str
		id = num
	}
	return param, paramName, id
}
