package parser

import (
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/driver"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	web "github.com/igor-koniukhov/webLogger/v3"
	"sync"
)

type RestMenuParserInterface interface {
	GetListSuppliers() (suppliers *model.Suppliers)
	GetListMenuItems(id int) (menu *model.Menu)
	ParsedDataWriter()
}

type RestMenuParser struct {
	App *config.AppConfig
}

var ParseRestMenu *RestMenuParser

func NewRestMenuParser(app *config.AppConfig) *RestMenuParser {
	return &RestMenuParser{App: app}
}
func NewRestMenu(r *RestMenuParser) {
	ParseRestMenu = r
}



var URL = "http://foodapi.true-tech.php.nixdev.co/restaurants"
var wg sync.WaitGroup

func (r *RestMenuParser) GetListSuppliers() (suppliers *model.Suppliers) {
	_ = json.Unmarshal(driver.GetBodyConnection(URL), &suppliers)
	return
}
func (r *RestMenuParser) GetListMenuItems(id int) (menu *model.Menu) {
	var URLMenu = fmt.Sprintf("%s/%v/menu", URL, id)
	_ = json.Unmarshal(driver.GetBodyConnection(URLMenu), &menu)
	return
}

func (r *RestMenuParser) ParsedDataWriter() {

	parsedSuppliers := r.GetListSuppliers()
	suppliersInDB, err := repository.Repo.SupplierRepositoryInterface.Create(parsedSuppliers)
	web.Log.Error(err, err)

	for _, restaurant := range suppliersInDB.Restaurants {
		menu := r.GetListMenuItems(restaurant.Id)
		id := <-r.App.ChanIdSupplier
		idSoftDel := id - len(suppliersInDB.Restaurants)
		repository.Repo.SupplierRepositoryInterface.SoftDelete(idSoftDel)
		for _, item := range menu.Items {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				repository.Repo.SupplierRepositoryInterface.SoftDelete(idSoftDel)
				_, _ = repository.Repo.ProductRepositoryInterface.Create(&item, id)
			}(id)
		}
		wg.Wait()
	}
}




