package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"io"
	"sync"

	"os"
)

var (
	err        error
	idSequence int
	DataUser   []*model.User
)
// Cut from main

/*userFileRepository := helpers.NewUserRepository(&app)
userCreateHandler := handlers.UserHandler{
	UserRepository: userFileRepository,
}*/

func CreateModel(modelName string, v interface{}) error {
	bytes, err := json.MarshalIndent(v, "", "   ")
	_ = checkReturnError(err)
	file, err := os.OpenFile(fmt.Sprintf("./datastore/%s.json", modelName), os.O_CREATE|os.O_RDWR, 0600)
	_ = checkReturnError(err)
	defer file.Close()

	//chunk of code for json write build structure
	fstat, err := file.Stat()
	_ = checkReturnError(err)
	fSize := fstat.Size()

	if fSize < 1 {
		_, err = file.WriteAt([]byte(`[`), 0)
		_ = checkReturnError(err)
		bytes = append(bytes, ']')
		_, err = file.WriteAt(bytes, fSize+1)
		_ = checkReturnError(err)
	}
	if fSize > 1 {
		bytes = append(bytes, ']')
		_, err = file.WriteAt([]byte(`, \n`), fSize-1)
		_, err = file.WriteAt(bytes, fSize+1)
		_ = checkReturnError(err)
	}
	return err
}

func checkReturnError(err error) error {
	if err != nil {
		err.Error()
	}
	return err
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


	err := CreateModel("users", user)
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

func (ufr Repository) Delete(id int) (*model.User, error) {
	var v *model.User
	for _, v := range DataUser {

		if v.ID == id {
			v.Status = "deleted"
			return v, err
		}
	}
	return v, err
}

func (ufr Repository) Update(u2 *model.User) *model.User {
	var v *model.User
	for _, v := range DataUser {
		if v.ID == u2.ID {
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

func (ufr *Repository) GetNextID() int {
	fl, err := os.OpenFile("./datastore/users.json", os.O_RDWR, 0600)
	repository.CheckErr(err)
	defer fl.Close()
	data, err := io.ReadAll(fl)
	err = json.Unmarshal(data, &DataUser)
	repository.CheckErr(err)
	idSequence = len(DataUser) - 1
	ufr.idMutex.Lock()
	idSequence += 1
	return idSequence
}