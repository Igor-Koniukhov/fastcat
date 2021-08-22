package model

type Product struct {
	ID           int32   `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	ImgReference string  `json:"img_reference"`
	Price        float64 `json:"price"`
}
