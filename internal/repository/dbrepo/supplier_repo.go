package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	web "github.com/igor-koniukhov/webLogger/v3"
	"sync"
	"time"
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
	DB *sql.DB
}

func NewSupplierRepository(app *config.AppConfig, DB *sql.DB) *SupplierRepo {
	return &SupplierRepo{App: app, DB: DB}
}
func (s SupplierRepo) Create(suppliers *model.Suppliers) (*model.Suppliers, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int
	stmtSql := fmt.Sprintf("INSERT INTO %s (name, image) VALUES (?, ?)", model.TabSuppliers)
	for _, restaurant := range suppliers.Restaurants {
		result, err := s.DB.ExecContext(ctx, stmtSql,
			restaurant.Name,
			restaurant.Image)
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
	return suppliers, nil
}

func (s SupplierRepo) Get(id int) *model.Supplier {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var supplier model.Supplier
	sqlStmt := fmt.Sprintf("SELECT id, name, image FROM %s WHERE id = ? ", model.TabSuppliers)
	fmt.Println(sqlStmt)
	err := s.DB.QueryRowContext(ctx, sqlStmt, id).Scan(
		&supplier.Id,
		&supplier.Name,
		&supplier.Image,
	)
	web.Log.Error(err, err)
	return &supplier
}

func (s SupplierRepo) GetAll() []model.Supplier {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var supplier model.Supplier
	var suppliers []model.Supplier
	sqlStmt := fmt.Sprintf("SELECT id, name, image FROM %s WHERE deleted_at IS NULL", model.TabSuppliers)
	stmt, err := s.DB.QueryContext(ctx, sqlStmt)
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=? ", model.TabSuppliers)
	_, err = s.DB.ExecContext(ctx, sqlStmt, id)
	web.Log.Error(err, err)
	return
}
func (s SupplierRepo) SoftDelete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("UPDATE %s SET deleted_at = ? WHERE id = ?", model.TabSuppliers)
	_, err := s.DB.ExecContext(ctx, sqlStmt, s.App.TimeFormat, id)
	web.Log.Error(err, err)
	return nil
}

func (s SupplierRepo) Update(id int, supplier *model.Supplier) *model.Supplier {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?, image=?, Name=?, Menu=? , WHERE id=?", model.TabSuppliers)
	_, err := s.DB.ExecContext(ctx, sqlStmt, supplier.Id, supplier.Image, supplier.Name, supplier.Menu, id)
	web.Log.Error(err, err)
	return supplier
}
