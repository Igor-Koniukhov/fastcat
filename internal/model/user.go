package model

const TableUser = "users"
type User struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	DeletedAt interface{} `json:"deleted_at"`
	CreatedAT string      `json:"created_at"`
	UpdatedAT string      `json:"updated_at"`
}



