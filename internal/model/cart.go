package model

type Cart struct {
	ID        int            `json:"id"`
	UserID    int            `json:"user_id"`
	ProductID int            `json:"product_id"`
	Items     []CartProducts `json:"items"`
	Status    string         `json:"status"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`

}
const TableCarts = "carts"

type CartProducts struct {
	ID        int `json:"id"`
	CartID    int `json:"cart_id"`
	ProductID int `json:"product_id"`
	Number    int `json:"number"`
}
