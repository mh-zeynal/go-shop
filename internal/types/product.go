package types

type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Off int `json:"off"`
}

func (p Product) TableName() string {
	return "products"
}

