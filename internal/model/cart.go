package model

type Cart struct{
	ID int32 `json:"id"`
	Items []Product `json:"items"`
}

