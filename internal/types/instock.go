package types

//a type that carries all available products in database
type InStock struct {
	AllProducts []Product `json:"all_products"`
}
