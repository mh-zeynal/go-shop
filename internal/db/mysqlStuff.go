package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shoppingServer/internal/types"
)

var db *gorm.DB

func IsConnectionStablished() bool {
	return db != nil
}

func MakeDatabase(user string, pass string, dbname string)  {
	dsn := user + ":" + pass + "@/" + dbname
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("database connection has failed")
	}
}

func GetAllProducts() types.InStock {
	temp := make([]types.Product, 4)
	db.Find(&temp)
	return types.InStock{AllProducts: temp}
}

func GetProducts(ids []int) *[]types.Product {
	selects := make([]types.Product, 0)
	for _, id := range ids{
		temp := types.Product{}
		_ = db.Where("id = ?", id).First(&temp)
		selects = append(selects, temp)
	}
	return &selects
}



