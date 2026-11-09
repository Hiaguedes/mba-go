package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	CategoryId int 
	Category Category
	SerialNumber SerialNumber
	gorm.Model
}

type Category struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Products []Product
}

type SerialNumber struct {
	ID int `gorm:"primaryKey"`
	Number string
	ProductID int
}

func main() {
	dsn := "test.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// db.Create(&[]Product{
	// 	{ Name: "NoBreak", Price: 429.90 },
	// 	{ Name: "Nvidia 5080", Price: 5550.00 },
	// })

	// var products []Product
	// db.Where("price > ?", 100).Find(&products)
	// db.Where("name LIKE ?", "%book%").Find(&products)

	// for _, product := range products {
	// 	fmt.Printf("Select: %+v", product)
	// }

	// var p Product


	// db.Model(&Product{}).Where("id = ?", 7).Update("name", "Nvidia RTX 5080")
	// db.Model(&Product{}).Delete("id = ?", 6)

	// category := Category{ Name: "Utilidades" }
	// db.Create(&category)

	// db.Create(&Product{
	// 	Name: "Vaso de Planta",
	// 	Price: 5.00,
	// 	CategoryId: category.ID,
	// })

	// db.Create(&SerialNumber{
	// 	Number: "12345",
	// 	ProductID: 10,
	// })

	var categories []Category

	db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories)
		for _, category := range categories {
		fmt.Printf("Select Categories: %+v\n", category)
	}

	fmt.Println("----------------")

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products) // preload associa produto a categoria

	for _, product := range products {
		fmt.Printf("Select Products: %+v\n", product)
	}
}