package repoForTest

import (
	"database/sql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/models"
)

type UserTestRepository interface {
	Create(user *models.User) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	GetAll() []models.User
	Delete(id int) error
	Update(id int, user *models.User) *models.User
	GetUserByEmail(email string) (*models.User, error)
}

type UserTestRepo struct {
	DB    *sql.DB
	Users []models.User
	User  models.User
	App *config.AppConfig
}

func NewUserTestRepository(app *config.AppConfig, DB *sql.DB) *UserTestRepo {
	return &UserTestRepo{ App: app, DB: DB}
}

func (usr UserTestRepo) Create(user *models.User) (*models.User, error) {

	return nil, nil
}

func (usr UserTestRepo) GetUserByID(id int) (*models.User, error) {
	return nil, nil
}

func (usr UserTestRepo) GetAll() []models.User {

	return nil
}

func (usr UserTestRepo) Delete(id int) error {

	return nil
}

func (usr UserTestRepo) Update(id int, user *models.User) *models.User {

	return nil
}

func (usr UserTestRepo) GetUserByEmail(email string) (*models.User, error) {

	return nil, nil
}


