package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"shoppingServer/internal/calculations"
	"shoppingServer/internal/db"
	"shoppingServer/internal/types"
)

func ShowGoods(c echo.Context) error {
	if !db.IsConnectionStablished() {
		db.MakeDatabase("root", "m@96@s97", "gamedatabase")
	}
	return c.JSON(http.StatusOK, db.GetAllProducts())
}

func ReceiveSelectedGoods(c echo.Context) error {
	if !db.IsConnectionStablished() {
		db.MakeDatabase("root", "m@96@s97", "gamedatabase")
	}
	temp := types.Basket{}
	err := json.NewDecoder(c.Request().Body).Decode(&temp)
	if err != nil {
		fmt.Println("no json passed to the server")
	}
	res := calculations.GetFinalPrice(*db.GetProducts(temp.SelectedGoods))
	receipt := types.Receipt{FinalPrice: res, SelectedGoods: *db.GetProducts(temp.SelectedGoods)}
	return c.JSON(http.StatusOK, receipt)
}
