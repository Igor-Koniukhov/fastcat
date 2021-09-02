package helpers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"io/ioutil"
	"log"
	"net/http"
)

type RestMenuRepositoryI interface {
	GetRestaurants()
	GetMenu()
}

type RestMenuRepository struct {
	App        *config.AppConfig
	Suppliers  *Suppliers
	Menu       *Menu
	Restaurant *Restaurant
}

var RepoRestMenu *RestMenuRepository

func NewRestMenuRepository(app *config.AppConfig) *RestMenuRepository {
	return &RestMenuRepository{App: app}
}
func NewRestMenu(r *RestMenuRepository) {
	RepoRestMenu = r
}

type Suppliers struct {
	Restaurants []Restaurant `json:"restaurants"`
}

type Restaurant struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Menu struct {
	Menu []Items `json:"menu"`
}

type Items struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Type  string  `json:"type"`
	//Ingredients []string`json:"ingredients"`
	Ingredients []string `json:"ingredients"`
}
type Ingredients struct {
	Ingredients string `json:"ingredients"`
}

var ctx context.Context

const TabSuppliers = "suppliers"
const TabMenu = "menu"

func (r RestMenuRepository) GetRestaurants() {
	var url string = "http://foodapi.true-tech.php.nixdev.co/restaurants"
	var suppliers Suppliers

	json.Unmarshal(responseConnection(url), &suppliers)

	tx, err := r.App.DB.Begin()
	if err != nil {
		fmt.Println(err)
	}
	tx.Exec("DELETE FROM " + TabSuppliers)
	tx.Exec("DELETE FROM " + TabMenu)
	stmtSQL, err := tx.Prepare("INSERT suppliers SET id=?, name=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmtSQL.Close()

	for _, restaurant := range suppliers.Restaurants {
		_, err = stmtSQL.Exec(
			restaurant.Id,
			restaurant.Name,
		)
r.getMenu(restaurant.Id, tx)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
	}
}

func (r RestMenuRepository) getMenu(id int, tx *sql.Tx) {
	var menu Menu
	url := fmt.Sprintf("http://foodapi.true-tech.php.nixdev.co/restaurants/%v/menu", id)
	json.Unmarshal(responseConnection(url), &menu)
	stmtSQl, err := tx.Prepare("INSERT INTO " + TabMenu + " (name, price, type, ingredients) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	defer stmtSQl.Close()
	for _, items := range menu.Menu {
		ingredients, err := json.MarshalIndent(items.Ingredients, "", "")
		if err != nil {
			fmt.Println(err)
		}
		_, err = stmtSQl.Exec(
			items.Name,
			items.Price,
			items.Type,
			ingredients,
		)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func responseConnection(url string) (response []byte) {
	conn, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Body.Close()
	response, err = ioutil.ReadAll(conn.Body)
	if err != nil {
		fmt.Println(err)
	}
	return
}
