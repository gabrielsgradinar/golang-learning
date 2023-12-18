package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "postgres://postgres:admin@localhost:5432/app?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// // create category
	// category := Category{Name: "Eletrônicos"}
	// db.Create(&category)

	// // create product
	// products := []Product{
	// 	{Name: "Notebook", Price: 1100.00, CategoryID: category.ID},
	// }
	// db.Create(&products)

	// db.Create(&SerialNumber{
	// 	Number: "123456",
	// 	ProductID: 1,
	// })

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name, ", Serial Number:", product.SerialNumber.Number)
		}
	}

}
