package models

const TableOrders = "orders"
const TableAddress = "address"

type Order struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Product   string `json:"product"`
	Price     string `json:"price"`
	UserName  string `json:"user"`
	UserEmail string `json:"user_email"`
	Phone     string `json:"phone"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Address struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	City        string `json:"city"`
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	FlatNumber  string `json:"flat_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
 type Ord struct {
 	Products []Item
 	User User
 }
type Prod struct {
	Title string `json:"title"`
	Price string `json:"price"`
}