package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
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

	// create category
	// category := Category{Name: "Eletrônicos"}
	// db.Create(&category)

	// create product
	// db.Create(&Product{
	// 	Name:       "Notebook",
	// 	Price:      1000.00,
	// 	CategoryID: category.ID,
	// })

	// db.Create(&Product{
	// 	Name:       "Mouse",
	// 	Price:      50.00,
	// 	CategoryID: category.ID,
	// })

	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products{
		fmt.Println(product.Name, product.Category.Name)
	}

}
