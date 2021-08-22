package model

type Supplier struct {
	ID          int    `json:"id"`
	Category    string `json:"category"`
	Name        string `json:"name"`
	Address     string `json:"addressName"`
	Description string `json:"description"`
	IconURL     string `json:"icon_url"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
