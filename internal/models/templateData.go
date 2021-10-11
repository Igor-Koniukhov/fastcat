package models

type TemplateData struct {
	Suppliers    []Supplier
	Supplier     *Supplier
	LoginRequest LoginRequest
	Products     []Item
	Cart         *Cart
	ErrorMessage string
	UserName     string
	StringMap    map[string]string
	IntMap       map[string]int
	FloatMap     map[string]float32
	Data         map[string]interface{}
	Flash        string
	Warning      string
	Error        string
}
