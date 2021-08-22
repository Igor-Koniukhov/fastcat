package model

type Product struct {
	ID           int     `json:"id"`
	SupplierID   int     `json:"supplier_id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	ImgReference string  `json:"img_reference"`
	Price        float64 `json:"price"`
	Status       string  `json:"status"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

