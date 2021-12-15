package main

import (
	"github.com/labstack/echo"
	"shoppingServer/internal/handlers"
)

func main() {
	e := echo.New()
	e.GET("/goods", handlers.ShowGoods)
	e.POST("/buy", handlers.ReceiveSelectedGoods)
	e.Logger.Fatal(e.Start(":8080"))
}

