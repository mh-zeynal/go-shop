package types

//a type that carries all products that user has selected
type Basket struct {
	SelectedGoods []int `json:"selected_goods"`
}
