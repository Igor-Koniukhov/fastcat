package main

import (
	"fmt"
	"github.com/igor-koniukhov/fastcat/model"
	r "github.com/igor-koniukhov/fastcat/repository"
)

func main() {
	u := model.User{
		Name:     "Name",
		Email:    "Name@gmail.com",
		Password: "password",
		Status:   "created",
	}
	u2 := model.User{
		ID:       34,
		Name:     "Name",
		Email:    "Name@gmail.com",
		Password: "password",
		Status:   "created",
	}
	email := "igor@gmail.com"
	var id int32 = 20

	userRepository := r.NewUserFileRepository()
	storedUser, err := userRepository.Create(&u)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	printEmail := userRepository.Get(&email)
	allUser := userRepository.GetAll()
	delete, err := userRepository.Delete(id)
	r.CheckErr(err)
	edit := userRepository.Edit(&u2)

	for _, data := range allUser {
		fmt.Println(*data, "this is from GetAll")
	}

	fmt.Println(storedUser)
	fmt.Println(printEmail)
	fmt.Println(delete)
	fmt.Println(edit)
}
