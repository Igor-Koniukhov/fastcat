package repoForTest

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
)

type CartTestRepository interface {
	Create(cart *models.Cart) (*models.Cart, error)
	Get(id int) *models.Cart
	GetAll() []models.Cart
	Delete(id int) error
	Update(id int, u *models.Cart) *models.Cart
}

type CartTestRepo struct {
	DB  *sql.DB
	App *config.AppConfig
}

func NewCartTestRepository(app *config.AppConfig, DB *sql.DB) *CartTestRepo {
	return &CartTestRepo{App: app, DB: DB}
}

func (c CartTestRepo) Create(cart *models.Cart) (*models.Cart, error) {

	return nil, nil
}

func (c CartTestRepo) Get(id int) *models.Cart {

	return nil
}

func (c CartTestRepo) GetAll() []models.Cart {

	return nil
}

func (c CartTestRepo) Delete(id int) error {

	return nil
}

func (c CartTestRepo) Update(id int, cart *models.Cart) *models.Cart {

	return nil
}
