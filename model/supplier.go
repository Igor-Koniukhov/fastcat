package model

type Supplier struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	IconURL     string `json:"icon_url"`
	AddressName string `json:"addressName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
