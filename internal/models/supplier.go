package models


type Suppliers struct {
	Restaurants []Supplier `json:"restaurants"`
}

type Supplier struct {
	Id   int    `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Menu []Item`json:"menu"`

}

const TabSuppliers = "suppliers"
