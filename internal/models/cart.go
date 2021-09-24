package models

const TableCarts = "carts"

type CartResponse struct {
	ID              int `json:"id"`
	User            []byte
	AddressDelivery string `json:"address_delivery"`
	CartBody        []byte `json:"order_body"`
	Amount          string `json:"amount"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
type Cart struct {
	ID              int `json:"id"`
	User            User
	AddressDelivery string `json:"address_delivery"`
	CartBodies        []CartBody
	Amount          string `json:"amount"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type CartBody struct {
	ProductID  string `json:"product_id"`
	SupplierID string `json:"supplier_id"`
	Title      string `json:"title"`
	Price      string `json:"price"`
}
