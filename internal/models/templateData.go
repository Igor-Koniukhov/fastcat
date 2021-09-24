package models

type TemplateData struct {
	Suppliers []Supplier
	Supplier  *Supplier
	Products  []Item
	Cart      *Cart
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	Flash     string
	Warning   string
	Error     string
}
