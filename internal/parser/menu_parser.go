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

func supplierAppConfigProvider(a *config.AppConfig) *repository.SupplierRepository {
	repo := repository.NewSupplierRepository(a)
	repository.NewRepoS(repo)
	return repo
}
func productAppConfigProvider(a *config.AppConfig) *repository.ProductRepository {
	repo := repository.NewProductRepository(a)
	repository.NewRepoP(repo)
	return repo
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
	supplierAppConfigProvider(r.App)
	productAppConfigProvider(r.App)

	parsedSuppliers := r.GetListSuppliers()
	suppliersInDB, err := repository.RepoS.CreateSupplier(parsedSuppliers)
	web.Log.Error(err)

	for _, restaurant := range suppliersInDB.Restaurants {
		menu := r.GetListMenuItems(restaurant.Id)
		id := <-r.App.ChanIdSupplier
		idSoftDel := id - len(parsedSuppliers.Restaurants)
		repository.RepoS.SoftDelete(idSoftDel)

		for _, item := range menu.Items {
			wg.Add(1)
			r.App.ChanMutex <- 1
			go func(id int) {
				defer wg.Done()
				repository.RepoP.SoftDelete(idSoftDel)
				_, _ = repository.RepoP.CreateProduct(&item, id)
			}(id)
			<-r.App.ChanMutex
		}
		wg.Wait()
	}
}




