package parser

import (
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/driver"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	web "github.com/igor-koniukhov/webLogger/v2"
	"time"

	"log"
	"sync"
)

type RestMenuParserInterface interface {
	GetListSuppliers() (suppliers *models.Suppliers, err error)
	GetListMenuItems(id int) (menu *models.Menu, err error)
	ParsedDataWriter() error
}

type RestMenuParser struct {
	App *config.AppConfig
	wg  sync.WaitGroup
}

var ParseRestMenu *RestMenuParser

func NewRestMenuParser(app *config.AppConfig) *RestMenuParser {
	return &RestMenuParser{App: app}
}
func NewRestMenu(r *RestMenuParser) {
	ParseRestMenu = r
}

func (r *RestMenuParser) GetListSuppliers() (suppliers *models.Suppliers, err error) {
	URL := "http://foodapi.true-tech.php.nixdev.co/suppliers"
	err = json.Unmarshal(driver.GetBodyConnection(URL), &suppliers)
	if err != nil {
		web.Log.Fatal(err)
		return suppliers, err
	}

	return suppliers, nil
}

func (r *RestMenuParser) GetListMenuItems(id int) (menu *models.Menu, err error) {
	URL := "http://foodapi.true-tech.php.nixdev.co/suppliers"
	var URLMenu = fmt.Sprintf("%s/%v/menu", URL, id)
	err = json.Unmarshal(driver.GetBodyConnection(URLMenu), &menu)
	if err != nil {
		web.Log.Fatal(err)
	}
	return menu, nil
}

func (r *RestMenuParser) ParsedDataWriter() error {
	parsedSuppliers, err := r.GetListSuppliers()
	ch := make(chan int, 1)
	defer close(ch)
	r.App.ChanIdSupplier = make(chan int, len(parsedSuppliers.Suppliers))
	defer close(r.App.ChanIdSupplier)
	if err != nil {
		web.Log.Fatal(err)
		return err
	}
	suppliersInDB, _, err := repository.Repo.SupplierRepository.Create(parsedSuppliers)
	if err != nil {
		web.Log.Fatal(err)
		return err
	}
	for _, restaurant := range suppliersInDB.Suppliers {
		menu, err := r.GetListMenuItems(restaurant.Id)
		if err != nil {
			web.Log.Info(err)
			log.Fatal(err)
			return err
		}
		id := <-r.App.ChanIdSupplier
		idSoftDel := id - len(suppliersInDB.Suppliers)
		err = repository.Repo.SupplierRepository.SoftDelete(idSoftDel)
		if err != nil {
			log.Fatal(err)
			return err
		}
		for _, item := range menu.Items {
			r.wg.Add(1)
			go func(id int) {
				defer r.wg.Done()
				err := repository.Repo.ProductRepository.SoftDelete(idSoftDel)
				if err != nil {
					web.Log.Fatal(err)
					return
				}
				_, err = repository.Repo.ProductRepository.Create(&item, id)
				if err != nil {
					web.Log.Fatal(err)
					return
				}
				ch <- 1
			}(id)
			<-ch
		}
		r.wg.Wait()
	}
	return nil
}

func RunUpToDateSuppliersInfo(t time.Duration) error {
	for {
		err := ParseRestMenu.ParsedDataWriter()
		if err != nil {
			web.Log.Fatal(err)
			return err
		}
		fmt.Println("Menu is up-to-date ", time.Now())
		time.Sleep(time.Second * t)
	}
	return nil
}
