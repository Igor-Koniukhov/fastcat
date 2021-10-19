package model

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Status      string `json:"status"`
	CreatedAT   string `json:"created_at"`
	UpdatedAT   string `json:"updated_at"`
}