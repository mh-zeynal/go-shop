package types

//a type that carries related data about a product
type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Off int `json:"off"`
}

/*this method makes query easier for gorm methods to recognize
that which table we are working with*/
func (p Product) TableName() string {
	return "products"
}

