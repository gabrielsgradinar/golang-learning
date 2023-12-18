package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	ID    int  `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "postgres://postgres:admin@localhost:5432/app?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// create
	// db.Create(&Product{
	// 	Name: "Notebook",
	// 	Price: 1000.00,
	// })

	// create batch
	// products := []Product{
	// 	{Name: "Notebook", Price: 1100.00},
	// 	{Name: "Mouse", Price: 50.00},
	// 	{Name: "Keyboard", Price: 200.00},
	// }
	// db.Create(&products)

	// select one
	// var product Product
	// db.First(&product, 1)
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// select all
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products{
	// 	fmt.Println(product)
	// }

	// where
	// var products []Product
	// db.Where("name LIKE ?", "%k%").Find(&products)
	// for _, product := range products{
	// 	fmt.Println(product)
	// }

	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Mouse"
	// db.Save(&p)
	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)
	db.Delete(&p2)

}