package repository

import (
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	web "github.com/igor-koniukhov/webLogger/v3"
	"sync"
)

var wg sync.WaitGroup

type SupplierRepository interface {
	Create(suppliers *model.Suppliers) (*model.Suppliers, error)
	Get(id int) *model.Supplier
	GetAll() []model.Supplier
	Delete(id int) error
	SoftDelete(id int) error
	Update(id int, supplier *model.Supplier) *model.Supplier
}

type SupplierRepo struct {
	App *config.AppConfig
}

func NewSupplierRepository(app *config.AppConfig) *SupplierRepo {
	return &SupplierRepo{App: app}
}
func (s SupplierRepo) Create(suppliers *model.Suppliers) (*model.Suppliers, error) {
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
			s.App.ChanIdSupplier <- id
			wg.Done()
		}(id)
	}
	wg.Wait()
	return suppliers, err
}

func (s SupplierRepo) Get(id int) *model.Supplier {
	var supplier model.Supplier
	sqlStmt := fmt.Sprintf("SELECT id, name, image FROM %s WHERE id = ? ", model.TabSuppliers)
	fmt.Println(sqlStmt)
	err := s.App.DB.QueryRow(sqlStmt, id).Scan(
		&supplier.Id,
		&supplier.Name,
		&supplier.Image,
	)
	web.Log.Error(err, err)
	return &supplier
}

func (s SupplierRepo) GetAll() []model.Supplier {
	var supplier model.Supplier
	var suppliers []model.Supplier
	sqlStmt := fmt.Sprintf("SELECT id, name, image FROM %s WHERE deleted_at IS NULL", model.TabSuppliers)
	stmt, err := s.App.DB.Query(sqlStmt)
	web.Log.Error(err, err)
	for stmt.Next() {
		_ = stmt.Scan(
			&supplier.Id,
			&supplier.Name,
			&supplier.Image,
		)
		suppliers = append(suppliers, supplier)
	}
	return suppliers
}

func (s SupplierRepo) Delete(id int) (err error) {
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=? ", model.TabSuppliers)
	_, err = s.App.DB.Exec(sqlStmt, id)
	fmt.Println(sqlStmt)
	web.Log.Error(err, err)
	return
}
func (s SupplierRepo) SoftDelete(id int) error {
	sqlStmt := fmt.Sprintf("UPDATE %s SET deleted_at = ? WHERE id = ?", model.TabSuppliers)
	_, err := s.App.DB.Exec(sqlStmt, s.App.TimeFormat, id)
	web.Log.Error(err, err)
	return nil
}

func (s SupplierRepo) Update(id int, supplier *model.Supplier) *model.Supplier {
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, image=?, Name=?, Menu=? , WHERE id=?", model.TabSuppliers)
	_, err := s.App.DB.Exec(sqlStmt, supplier.Id, supplier.Image, supplier.Name, supplier.Menu, id)
	web.Log.Error(err, err)
	return supplier
}
