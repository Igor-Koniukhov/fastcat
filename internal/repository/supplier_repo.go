package repository

import (
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	web "github.com/igor-koniukhov/webLogger/v2"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type SupplierRepositoryI interface {
	CreateSupplier(u *model.Supplier) (*model.Supplier, error)
	GetSupplier(nameParam, param *string, ) *model.Supplier
	GetAllUSuppliers() *[]model.Supplier
	DeleteSupplier(id int) error
	UpdateSupplier(id int, u *model.Supplier) *model.Supplier
	Param(r *http.Request) (string, string, int)
}

var wg sync.WaitGroup
var RepoS *SupplierRepository
var supplier model.Supplier

type SupplierRepository struct {
	App *config.AppConfig
}

func NewSupplierRepository(app *config.AppConfig) *SupplierRepository {
	return &SupplierRepository{App: app}
}

func NewRepoS(r *SupplierRepository) {
	RepoS = r
}

func (s SupplierRepository) CreateSupplier(suppliers *model.Suppliers) (*model.Suppliers, error) {
	var id int
	stmtSql := fmt.Sprintf("INSERT INTO %s (name, image) VALUES (?, ?)", model.TabSuppliers)

	stmt, err := s.App.DB.Prepare(stmtSql)
	defer stmt.Close()

	for _, restaurant := range suppliers.Restaurants {
		result, err := stmt.Exec(restaurant.Name, restaurant.Image)
		web.Log.Error(err)
		lastInsertedID, err := result.LastInsertId()
		web.Log.Error(err)
		id = int(lastInsertedID)
		wg.Add(1)
		go func(id int) {
			s.App.ChanelSupplierId <- id
			wg.Done()
		}(id)
	}
	wg.Wait()
	return suppliers, err
}

func (s SupplierRepository) GetSupplier(nameParam, param *string) *model.Supplier {
	sqlStmt := fmt.Sprintf("SELECT id, name, image FROM %s WHERE %s = '%s' ", model.TabSuppliers, *nameParam, *param)
	fmt.Println(sqlStmt)
	err := s.App.DB.QueryRow(sqlStmt).Scan(
		&supplier.Id,
		&supplier.Name,
		&supplier.Image,
	)
	CheckErr(err)
	return &supplier
}

func (s SupplierRepository) GetAllUSuppliers() *[]model.Supplier {
	var supplier model.Supplier
	var suppliers []model.Supplier
	sqlStmt := fmt.Sprintf("SELECT id, name, image FROM %s ", model.TabSuppliers)
	stmt, err := s.App.DB.Query(sqlStmt)
	web.Log.Error(err, err)
	for stmt.Next() {
		stmt.Scan(
			&supplier.Id,
			&supplier.Name,
			&supplier.Image,
		)
		suppliers = append(suppliers, supplier)
	}
	return &suppliers
}

func (s SupplierRepository) DeleteSupplier(id int) (err error) {
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=? ", model.TabSuppliers)
	_, err = s.App.DB.Exec(sqlStmt, id)
	fmt.Println(sqlStmt)
	web.Log.Error(err, err)
	return
}
func (s SupplierRepository) SoftDelete(id int) error {
	sqlStmt := fmt.Sprintf("UPDATE %s SET deleted_at = ? WHERE id = ?", model.TabSuppliers)
	_, err := s.App.DB.Exec(sqlStmt, s.App.TimeFormat, id)
	web.Log.Error(err, err)
	return nil
}

func (s SupplierRepository) UpdateSupplier(id int, supplier *model.Supplier, ) *model.Supplier {
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, image=?, Name=?, Menu=? , WHERE id=?", model.TabSuppliers)
	_, err := s.App.DB.Exec(sqlStmt, supplier.Id, supplier.Image, supplier.Name, supplier.Menu, id)
	web.Log.Error(err, err)
	return supplier
}

func (s SupplierRepository) Param(r *http.Request) (string, string, int) {
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
