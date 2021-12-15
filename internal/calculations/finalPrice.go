package calculations

import "shoppingServer/internal/types"

func GetFinalPrice(products []types.Product) int {
	res := 0
	for _, v := range products{
		res += v.Price * (100 - v.Off) / 100
	}
	return res
}
