package repoForTest

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
)

type ProductTestRepository interface {
	Create(item *models.Item, id int) (*models.Item, error)
	Get(id int) *models.Item
	GetAll() []models.Item
	Delete(id int) (err error)
	SoftDelete(id int) error
	Update(id int, item *models.Item, ) *models.Item
}

type ProductTestRepo struct {
	DB *sql.DB
	App *config.AppConfig
}

func NewProductTestRepository(app *config.AppConfig, DB *sql.DB) *ProductTestRepo {
	return &ProductTestRepo{App: app, DB: DB}
}

func (p ProductTestRepo) Create(item *models.Item, id int) (*models.Item, error) {

	return nil, nil
}

func (p ProductTestRepo) Get(id int) *models.Item {

	return nil
}

func (p ProductTestRepo) GetAll() []models.Item {

	return nil
}

func (p ProductTestRepo) Delete(id int) (err error) {

	return nil
}

func (p ProductTestRepo) SoftDelete(id int) error {

	return nil
}

func (p ProductTestRepo) Update(id int, item *models.Item) *models.Item {

	return nil
}
