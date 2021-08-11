package repository

import (
	"encoding/json"
	"github.com/igor-koniukhov/fastcat/helpers"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"io"
	"log"
	"os"
	"sync"
)

var (
	err        error
	idSequence int32
	DataUser   []*model.User

)

type UserRepository interface {
	Create(u *model.User) (*model.User, error)
	Get(email *string) *model.User
	GetAll() []*model.User
	Delete(id int32) (*model.User, error)
	Edit(id int32, u *model.User) *model.User
}

type UserDBRepository struct {

}

func (u2 UserDBRepository) Create(u *model.User) (*model.User, error) {
	panic("implement me")
}

func (u2 UserDBRepository) Get(email *string) *model.User {
	panic("implement me")
}

func (u2 UserDBRepository) GetAll() []*model.User {
	panic("implement me")
}

func (u2 UserDBRepository) Delete(id int32) (*model.User, error) {
	panic("implement me")
}

func (u2 UserDBRepository) Edit(id int32, u *model.User) *model.User {
	panic("implement me")
}



type Repository struct {
	App *config.AppConfig
	idMutex *sync.Mutex
}

func NewUserRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App : a,
		idMutex: &sync.Mutex{},
	}
}

func (ufr Repository) Create(user *model.User) (*model.User, error) {
	user.ID = ufr.GetNextID()


	err := helpers.CreateModel("users", user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ufr Repository) Get(email *string) *model.User {
	var v *model.User
	for _, v := range DataUser {
		if v.Email == *email {
			return v
		}
	}
	return v
}

func (ufr Repository) GetAll() []*model.User {
	return DataUser
}

func (ufr Repository) Delete(id int32) (*model.User, error) {
	var v *model.User
	for _, v := range DataUser {

		if v.ID == id {
			v.Status = "deleted"
			return v, err
		}
	}
	return v, err
}

func (ufr Repository) Edit(id int32, u2 *model.User) *model.User {
	var v *model.User
	for _, v := range DataUser {
		if v.ID == id {
			v.ID = u2.ID
			v.Name = u2.Name
			v.Email = u2.Email
			v.PhoneNumber = u2.PhoneNumber
			v.Password = u2.Password
			v.Status = u2.Status

			return v
		}
	}
	return v
}

func (ufr *Repository) GetNextID() int32 {
	fl, err := os.OpenFile("./datastore/users.json", os.O_RDWR, 0600)
	CheckErr(err)
	defer fl.Close()
	data, err := io.ReadAll(fl)
	err = json.Unmarshal(data, &DataUser)
	CheckErr(err)
	idSequence = int32(len(DataUser) - 1)
	ufr.idMutex.Lock()
	idSequence += 1
	return idSequence
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
