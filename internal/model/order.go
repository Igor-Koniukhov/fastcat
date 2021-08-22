package model

type Order struct {
	ID        int32  `json:"id"`
	UserID    int    `json:"user_id"`
	CartID    int    `json:"cart_id"`
	AddressID int    `json:"address_id"`
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

