package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"log"
	"sync"
	"time"
)



type SupplierRepository interface {
	Create(suppliers *models.Suppliers) (*models.Suppliers, int, error)
	Get(id int) *models.Supplier
	GetAll() []models.Supplier
	GetAllBySchedule(start, end string) []models.Supplier
	Delete(id int) error
	SoftDelete(id int) error
	Update(id int, supplier *models.Supplier) *models.Supplier
}

type SupplierRepo struct {
	DB  *sql.DB
	App *config.AppConfig
	wg sync.WaitGroup
}

func NewSupplierRepository(app *config.AppConfig, DB *sql.DB) *SupplierRepo {
	return &SupplierRepo{App: app, DB: DB}
}
func (s SupplierRepo) Create(suppliers *models.Suppliers) (*models.Suppliers, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int
	var err error
	stmtSql := fmt.Sprintf("INSERT INTO %s (name, type, image, opening, closing) VALUES (?, ?, ?, ?, ?)", models.TabSuppliers)
	for _, restaurant := range suppliers.Suppliers {
		result, err := s.DB.ExecContext(ctx, stmtSql,
			restaurant.Name,
			restaurant.Type,
			restaurant.Image,
			restaurant.WorkingHours.Opening,
			restaurant.WorkingHours.Closing,
			)
		if err != nil {
			log.Println(err)
		}
		lastInsertedID, err := result.LastInsertId()
		if err != nil {
			log.Println(err)
		}
		id = int(lastInsertedID)

		s.wg.Add(1)
		go func(id int) {
			s.App.ChanIdSupplier <- id
			s.wg.Done()
		}(id)
	}
	s.wg.Wait()

	return suppliers, id, err
}

func (s SupplierRepo) Get(id int) *models.Supplier {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var supplier models.Supplier
	sqlStmt := fmt.Sprintf("SELECT id, name, type, image, opening, closing FROM %s WHERE id = ? ", models.TabSuppliers)
	err := s.DB.QueryRowContext(ctx, sqlStmt, id).Scan(
		&supplier.Id,
		&supplier.Name,
		&supplier.Type,
		&supplier.Image,
		&supplier.WorkingHours.Opening,
		&supplier.WorkingHours.Closing,
	)
	if err != nil {
		log.Println(err)
	}

	return &supplier
}

func (s SupplierRepo) GetAll() []models.Supplier {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var supplier models.Supplier
	var suppliers []models.Supplier
	sqlStmt := fmt.Sprintf("SELECT id, name, type, image, opening, closing FROM %s WHERE deleted_at IS NULL", models.TabSuppliers)
	stmt, err := s.DB.QueryContext(ctx, sqlStmt)
	if err != nil {
		log.Println(err)
	}
	for stmt.Next() {
		_ = stmt.Scan(
			&supplier.Id,
			&supplier.Name,
			&supplier.Type,
			&supplier.Image,
			&supplier.WorkingHours.Opening,
			&supplier.WorkingHours.Closing,
		)
		suppliers = append(suppliers, supplier)
	}

	return suppliers
}

func (s SupplierRepo) GetAllBySchedule(start, end string) []models.Supplier {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var supplier models.Supplier
	var suppliers []models.Supplier
	sqlStmt := fmt.Sprintf("SELECT id, name, type, image, opening, closing FROM %s WHERE deleted_at IS NULL AND opening='%s' AND closing='%s' ", models.TabSuppliers, start, end)
	stmt, err := s.DB.QueryContext(ctx, sqlStmt)
	if err != nil {
		log.Println(err)
	}
	for stmt.Next() {
		_ = stmt.Scan(
			&supplier.Id,
			&supplier.Name,
			&supplier.Type,
			&supplier.Image,
			&supplier.WorkingHours.Opening,
			&supplier.WorkingHours.Closing,
		)
		suppliers = append(suppliers, supplier)
	}

	return suppliers
}

func (s SupplierRepo) Delete(id int) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE id=? ", models.TabSuppliers)
	_, err = s.DB.ExecContext(ctx, sqlStmt, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return
}
func (s SupplierRepo) SoftDelete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	TimeFormat := time.Now().UTC().Format("2006-01-02 15:04:05.999999")
	sqlStmt := fmt.Sprintf("UPDATE %s SET deleted_at = ? WHERE id = ?", models.TabSuppliers)
	_, err := s.DB.ExecContext(ctx, sqlStmt, TimeFormat, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s SupplierRepo) Update(id int, supplier *models.Supplier) *models.Supplier {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("UPDATE %s SET id=?,name=?, type=?, image=?,  opening=?, closing=? , WHERE id=?", models.TabSuppliers)
	_, err := s.DB.ExecContext(ctx, sqlStmt,
		supplier.Id,
		supplier.Name,
		supplier.Type,
		supplier.Image,
		supplier.WorkingHours.Opening,
		supplier.WorkingHours.Closing, id)
	if err != nil {
		log.Println(err)
	}
	return supplier
}
