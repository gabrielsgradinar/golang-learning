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
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct{
	ID	int `gorm:"primaryKey"`
	Number string
	ProductID int
}

func main() {
	dsn := "postgres://postgres:admin@localhost:5432/app?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// create category
	category := Category{Name: "Eletrônicos"}
	db.Create(&category)

	// create product
	db.Create(&Product{
		Name:       "Notebook",
		Price:      1000.00,
		CategoryID: category.ID,
	})

	// create serial number
	db.Create(&SerialNumber{
		Number: "123456",
		ProductID: 1,
	})

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products{
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}

}
