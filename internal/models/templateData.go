package models

import "github.com/igor-koniukhov/fastcat/internal/forms"

type TemplateData struct {
	Suppliers       []Supplier
	Supplier        *Supplier
	LoginRequest    LoginRequest
	Products        []Item
	Cart            *Cart
	UserCabinetInfo []Cart
	Form            *forms.Form
	StringArr       []string
	ErrorMessage    string
	UserName        string
	StringMap       map[string]string
	StringSliceMap  map[string][]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	Flash           string
	Warning         string
	Error           string
}
