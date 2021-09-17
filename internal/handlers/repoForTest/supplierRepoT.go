package repoForTest

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"sync"
)

var wg sync.WaitGroup

type SupplierTestRepository interface {
	Create(suppliers *models.Suppliers) (*models.Suppliers, int, error)
	Get(id int) *models.Supplier
	GetAll() []models.Supplier
	Delete(id int) error
	SoftDelete(id int) error
	Update(id int, supplier *models.Supplier) *models.Supplier
}

type SupplierTestRepo struct {
	DB  *sql.DB
	App *config.AppConfig
}

func NewSupplierTestRepository(app *config.AppConfig, DB *sql.DB) *SupplierTestRepo {
	return &SupplierTestRepo{App: app, DB: DB}
}
func (s SupplierTestRepo) Create(suppliers *models.Suppliers) (*models.Suppliers, int, error) {

	return nil, 0, nil
}

func (s SupplierTestRepo) Get(id int) *models.Supplier {

	return nil
}

func (s SupplierTestRepo) GetAll() []models.Supplier {

	return nil
}

func (s SupplierTestRepo) Delete(id int) (err error) {

	return
}
func (s SupplierTestRepo) SoftDelete(id int) error {

	return nil
}

func (s SupplierTestRepo) Update(id int, supplier *models.Supplier) *models.Supplier {
	return nil
}
