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
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "postgres://postgres:admin@localhost:5432/app?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// // create category
	// category := Category{Name: "Eletrônicos"}
	// db.Create(&category)

	// // create product
	// products := []Product{
	// 	{Name: "Notebook", Price: 1100.00, CategoryID: category.ID},
	// 	{Name: "Mouse", Price: 50.00, CategoryID: category.ID},
	// 	{Name: "Keyboard", Price: 200.00, CategoryID: category.ID},
	// }
	// db.Create(&products)

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name)
		}
	}

}
