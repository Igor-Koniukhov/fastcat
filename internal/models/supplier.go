package models

type Suppliers struct {
	Suppliers []Supplier `json:"suppliers"`
}

type WorkingHours struct {
	Opening string `json:"opening"`
	Closing string `json:"closing"`
}

type Supplier struct {
	Id           int          `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	Image        string       `json:"image"`
	WorkingHours WorkingHours
}

const TabSuppliers = "suppliers"
