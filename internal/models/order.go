package models

const TableOrders = "orders"

type Order struct {
	ID         int       `json:"id"`
	SupplierId int       `json:"supplier_id"`
	BodyOrder  BodyOrder `json:"body_order"`
}

type BodyOrder struct {
	ID        int    `json:"id"`
	CartId    int    `json:"cart_id"`
	Title     string `json:"title"`
	Quantity  int    `json:"quantity"`
	Price     string `json:"price"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
