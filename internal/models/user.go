package models

const TableUsers = "users"
const Sessions = "sessions"

type User struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	Phone     string      `json:"phone"`
	Password  string      `json:"password"`
	DeletedAt interface{} `json:"deleted_at"`
	CreatedAT string      `json:"created_at"`
	UpdatedAT string      `json:"updated_at"`
	Session *UsersSessions
}

type UsersSessions struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Session string `json:"session"`
}
