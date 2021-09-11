package parser

import (
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/driver"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	web "github.com/igor-koniukhov/webLogger/v3"
	"sync"
)

type RestMenuParserInterface interface {
	GetListSuppliers() (suppliers *models.Suppliers)
	GetListMenuItems(id int) (menu *models.Menu)
	ParsedDataWriter()
}

type RestMenuParser struct {
	App *config.AppConfig
	wg sync.WaitGroup
}

var ParseRestMenu *RestMenuParser

func NewRestMenuParser(app *config.AppConfig) *RestMenuParser {
	return &RestMenuParser{App: app}
}
func NewRestMenu(r *RestMenuParser) {
	ParseRestMenu = r
}

func (r *RestMenuParser) GetListSuppliers() (suppliers *models.Suppliers) {
	URL := "http://foodapi.true-tech.php.nixdev.co/restaurants"
	_ = json.Unmarshal(driver.GetBodyConnection(URL), &suppliers)
	return
}
func (r *RestMenuParser) GetListMenuItems(id int) (menu *models.Menu) {
	URL := "http://foodapi.true-tech.php.nixdev.co/restaurants"
	var URLMenu = fmt.Sprintf("%s/%v/menu", URL, id)
	_ = json.Unmarshal(driver.GetBodyConnection(URLMenu), &menu)
	return
}

func (r *RestMenuParser) ParsedDataWriter() {

	parsedSuppliers := r.GetListSuppliers()
	suppliersInDB, err := repository.Repo.SupplierRepository.Create(parsedSuppliers)
	web.Log.Error(err, err)

	for _, restaurant := range suppliersInDB.Restaurants {
		menu := r.GetListMenuItems(restaurant.Id)
		id := <-r.App.ChanIdSupplier
		idSoftDel := id - len(suppliersInDB.Restaurants)
		repository.Repo.SupplierRepository.SoftDelete(idSoftDel)
		for _, item := range menu.Items {
			r.wg.Add(1)
			go func(id int) {
				defer r.wg.Done()
				repository.Repo.ProductRepository.SoftDelete(idSoftDel)
				_, _ = repository.Repo.ProductRepository.Create(&item, id)
			}(id)
		}
		r.wg.Wait()
	}
}




