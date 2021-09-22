package models

const TableCarts = "carts"



type CartResponse struct {
	ID              int `json:"id"`
	User            User
	AddressDelivery string `json:"address_delivery"`
	OrderBody       string `json:"order_body"`
	Amount          string `json:"amount"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}