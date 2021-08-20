package model

type Supplier struct {
	ID          int32  `json:"id"`
	Category    string `json:"category"`
	Name        string `json:"name"`
	AddressName string `json:"addressName"`
	Description string `json:"description"`
	IconURL     string `json:"icon_url"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
