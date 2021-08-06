package main

import (
	"fmt"
	"github.com/igor-koniukhov/fastcat/model"
	r "github.com/igor-koniukhov/fastcat/repository"
)



func main() {
	userRepository := r.NewUserFileRepository()

	u := model.User{
		Name:     "Name",
		Email:    "Name@gmail.com",
		PhoneNumber: "555 555 555",
		Password: "password",
		Status:   "created",
	}
	//u2 struct for imitation changing file
	u2 := model.User{
		ID:       34,
		Name:     "Arnold",
		Email:    "arni@gmail.com",
		PhoneNumber: "777 555 555",
		Password: "password",
		Status:   "created",
	}
	// email, id - var's for search and delete elements of file
	email := "igor@gmail.com"
	var id int32 = 20


	storedUser, err := userRepository.Create(&u)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	printEmail := userRepository.Get(&email)
	allUser := userRepository.GetAll()
	edit := userRepository.Edit(id, &u2)
	delete, err := userRepository.Delete(id)
	r.CheckErr(err)


	for _, data := range allUser {
		fmt.Println(*data, "this is from GetAll")
	}

	fmt.Println(storedUser)
	fmt.Println(printEmail)
	fmt.Println(delete)
	fmt.Println(edit)
}
