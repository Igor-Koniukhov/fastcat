package repoForTest

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
)

type OrderTestRepository interface {
	Create(or *models.Order) (*models.Order, error)
	Get(id int) *models.Order
	GetAll() *[]models.Order
	Delete(id int) error
	Update(id int, ord *models.Order) *models.Order
}

type OrderTestRepo struct {
	DB *sql.DB
	App *config.AppConfig
}

func NewOrderTestRepository(app *config.AppConfig, DB *sql.DB) *OrderTestRepo {
	return &OrderTestRepo{App: app, DB: DB}
}

func (o OrderTestRepo) Create(or *models.Order) (*models.Order, error) {

	return nil, nil
}

func (o OrderTestRepo) Get(id int) *models.Order {

	return nil
}

func (o OrderTestRepo) GetAll() *[]models.Order {

	return nil
}

func (o OrderTestRepo) Delete(id int) error {

	return nil
}

func (o OrderTestRepo) Update(id int, ord *models.Order) *models.Order {

	return nil
}
