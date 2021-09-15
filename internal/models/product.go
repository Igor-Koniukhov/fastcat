package models

const TabItems = "items"
const TabSuppliersItems = "suppliers_items"

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Type        string  `json:"type"`
	Ingredients []byte  `json:"ingredients"`
	SuppliersID int      `json:"suppliers_id"`
}

type Menu struct {
	Items []Item `json:"menu"`
}

type Item struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"`
	Type        string   `json:"type"`
	Ingredients []string `json:"ingredients"`
	SuppliersID int      `json:"suppliers_id"`
}

type RestaurantItems struct {
	Id         int `json:"id"`
	SupplierID int `json:"supplier_id"`
	ItemsId    int `json:"item_id"`
}
