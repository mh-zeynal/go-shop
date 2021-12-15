package types

type Receipt struct {
	FinalPrice int `json:"final_price"`
	SelectedGoods []Product `json:"selected_goods"`
}
