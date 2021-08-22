package model

type Order struct {
	ID           int32 `json:"id"`
	UserLocation []Address
}

type Address struct {
	City       string `json:"city"`
	Street     string `json:"street"`
	HoseNumber string `json:"hose_number"`
	FlatNumber string `json:"flat_number"`
}
