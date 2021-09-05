package parser

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	web "github.com/igor-koniukhov/webLogger/v2"
	"io/ioutil"
	"net/http"
	"sync"
)

type RestMenuParserInterface interface {
	ParseRestaurants()
	parseMenu(id int, tx *sql.Tx)
	responseBodyConnection(url string) (response []byte)
}

type RestMenuParser struct {
	App *config.AppConfig
}

var ParserRestMenu *RestMenuParser

func NewRestMenuRepository(app *config.AppConfig) *RestMenuParser {
	return &RestMenuParser{App: app}
}
func NewRestMenu(r *RestMenuParser) {
	ParserRestMenu = r
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

func (r *RestMenuParser) ParsedDataWriter() {
	var id int
	parsedSuppliers := r.GetListSuppliers()

	supplierAppConfigProvider(r.App)
	suppliersInDB, err := repository.RepoS.CreateSupplier(parsedSuppliers)
	web.Log.Error(err)
	productAppConfigProvider(r.App)
	for _, restaurant := range suppliersInDB.Restaurants {
		menu := r.GetListMenuItems(restaurant.Id)
		id = <-r.App.ChanelSupplierId
		for _, item := range menu.Items {
			wg.Add(1)
			r.App.ChanelLockUnlock <- 1
			go func(id int) {

				defer wg.Done()
				repository.RepoP.CreateProduct(&item, id)
			}(id)
			<-r.App.ChanelLockUnlock
		}
		wg.Wait()
	}
}

func (r *RestMenuParser) GetListSuppliers() (suppliers *model.Suppliers) {
	json.Unmarshal(r.responseBodyConnection(URL), &suppliers)
	return
}
func (r *RestMenuParser) GetListMenuItems(id int) (menu *model.Menu) {
	var URLMenu = fmt.Sprintf("%s/%v/menu", URL, id)
	json.Unmarshal(r.responseBodyConnection(URLMenu), &menu)
	return
}

func (r *RestMenuParser) responseBodyConnection(url string) (response []byte) {
	conn, err := http.Get(url)
	web.Log.Error(err)
	defer conn.Body.Close()
	response, err = ioutil.ReadAll(conn.Body)
	web.Log.Error(err)
	return
}
