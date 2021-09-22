package repoForTest

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
)

type OrderTestRepository interface {
	Create(or *models.Cart) (*models.Cart, error)
	Get(id int) *models.Cart
	GetAll() *[]models.Cart
	Delete(id int) error
	Update(id int, ord *models.Cart) *models.Cart
}

type OrderTestRepo struct {
	DB *sql.DB
	App *config.AppConfig
}

func NewOrderTestRepository(app *config.AppConfig, DB *sql.DB) *OrderTestRepo {
	return &OrderTestRepo{App: app, DB: DB}
}

func (o OrderTestRepo) Create(or *models.Cart) (*models.Cart, error) {

	return nil, nil
}

func (o OrderTestRepo) Get(id int) *models.Cart {

	return nil
}

func (o OrderTestRepo) GetAll() *[]models.Cart {

	return nil
}

func (o OrderTestRepo) Delete(id int) error {

	return nil
}

func (o OrderTestRepo) Update(id int, ord *models.Cart) *models.Cart {

	return nil
}
