package types

//this type contains final payable price for user and all products in his basket
type Receipt struct {
	FinalPrice int `json:"final_price"`
	SelectedGoods []Product `json:"selected_goods"`
}
