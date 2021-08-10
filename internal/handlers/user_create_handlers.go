package handlers

import (
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"github.com/igor-koniukhov/fastcat/internal/repository"
)

type UserHandler struct {
	UserRepository repository.UserRepository
}

func NewUserHandler(userRepository repository.UserRepository) *UserHandler {
	return &UserHandler{UserRepository: userRepository}
}

func (uch UserHandler) Handle() error {
	u := model.User{
		Name:        "Name",
		Email:       "Name@gmail.com",
		PhoneNumber: "555 555 555",
		Password:    "password",
		Status:      "created",
	}
	//u2 struct for imitation changing file
	u2 := model.User{
		ID:          34,
		Name:        "Arnold",
		Email:       "arni@gmail.com",
		PhoneNumber: "777 555 555",
		Password:    "password",
		Status:      "created",
	}
	// email, id - var's for search and delete elements of file
	email := "igor@gmail.com"
	var id int32 = 30

	storedUser, err := uch.UserRepository.Create(&u)
	if err != nil {
		fmt.Println(err.Error())
	}

	printEmail := uch.UserRepository.Get(&email)
	allUser := uch.UserRepository.GetAll()
	edit := uch.UserRepository.Edit(id, &u2)
	del, err := uch.UserRepository.Delete(id)
	repository.CheckErr(err)

	for _, data := range allUser {
		fmt.Println(*data, "this is from GetAll")
	}

	fmt.Println(storedUser)
	fmt.Println(printEmail)
	fmt.Println(del, " this is delete")
	fmt.Println(edit)


	return err
}
