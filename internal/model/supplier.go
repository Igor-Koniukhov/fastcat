package model

/*type Restaurants struct {
	Id          int    `json:"id"`
	Category    string `json:"category"`
	Name        string `json:"name"`
	Address     string `json:"addressName"`
	Description string `json:"description"`
	IconURL     string `json:"icon_url"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}*/
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
