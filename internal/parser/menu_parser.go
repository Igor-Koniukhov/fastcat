package parser

import (
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/driver"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"log"
	"sync"
	"time"
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
	URL := "http://foodapi.true-tech.php.nixdev.co/restaurants"
	err = json.Unmarshal(driver.GetBodyConnection(URL), &suppliers)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return suppliers, nil
}

func (r *RestMenuParser) GetListMenuItems(id int) (menu *models.Menu, err error) {
	URL := "http://foodapi.true-tech.php.nixdev.co/restaurants"
	var URLMenu = fmt.Sprintf("%s/%v/menu", URL, id)
	err = json.Unmarshal(driver.GetBodyConnection(URLMenu), &menu)
	if err != nil {
		log.Fatal(err)
	}
	return menu, nil
}

func (r *RestMenuParser) ParsedDataWriter() error {
	ch := make(chan int, 1)
	defer close(ch)
	r.App.ChanIdSupplier = make(chan int,4 )
	defer close(r.App.ChanIdSupplier)
	parsedSuppliers, err := r.GetListSuppliers()
	if err != nil {
		log.Fatal(err)
		return err
	}
	suppliersInDB, _, err := repository.Repo.SupplierRepository.Create(parsedSuppliers)
	if err != nil {
		log.Println(err)
		return err
	}
	for _, restaurant := range suppliersInDB.Restaurants {
		menu, err := r.GetListMenuItems(restaurant.Id)
		if err != nil {
			log.Fatal(err)
			return err
		}
		id := <-r.App.ChanIdSupplier
		idSoftDel := id - len(suppliersInDB.Restaurants)
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
					log.Println(err)
					return
				}
				_, err = repository.Repo.ProductRepository.Create(&item, id)
				if err != nil {
					log.Println(err)
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
		if err !=nil {
			log.Fatal(err)
			return err
		}
		fmt.Println("Menu is up-to-date ", time.Now())
		time.Sleep(time.Second * t)
	}
	return nil
}