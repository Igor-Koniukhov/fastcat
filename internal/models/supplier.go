package models

type Suppliers struct {
	Suppliers []Supplier `json:"suppliers"`
}

type WorkingHours struct {
	Closing string `json:"closing"`
	Opening string `json:"opening"`
}

type Supplier struct {
	Id           int          `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	Image        string       `json:"image"`
	WorkingHours WorkingHours
}

const TabSuppliers = "suppliers"
