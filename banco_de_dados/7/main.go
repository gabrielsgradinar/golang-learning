package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
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

	// category2 := Category{Name: "Cozinha"}
	// db.Create(&category2)

	// // create product
	// product := Product{
	// 	Name: "Panela Elétrica", Price: 300.00, Categories: []Category{category, category2},
	// }
	// db.Create(product)

	var products []Product
	err = db.Model(&Product{}).Preload("Categories").Find(&products).Error
	if err != nil {
		panic(err)
	}

	for _, product := range products {
		fmt.Println(product.Name, ":")
		for _, category := range product.Categories {
			fmt.Println("- ", category.Name)
		}
	}

}
