package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primary_key"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primary_key"`
	Name         string
	Price        int
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primary_key"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})
	if err != nil {
		panic(err)
	}

	// create category
	category := Category{Name: "Alimentos"}
	db.Create(&category)
	//
	//// create product
	product := Product{Name: "Rice", Price: 100, CategoryID: 1}
	db.Create(&product)

	// create serial number
	serialNumber := SerialNumber{Number: "123456", ProductID: 1}
	db.Create(&serialNumber)

	// select product
	//var products []Product
	//db.Preload("Category").Preload("SerialNumber").Find(&products)
	//for _, product := range products {
	//	fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	//}

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.Name, category.Name)
		}
	}

	// create one
	//db.Create(&Product{ID: id, Name: "Product 1", Price: 100})

	// create many
	//products := []Product{
	//	{ID: uuid.New().String(), Name: "Product 2", Price: 200},
	//	{ID: uuid.New().String(), Name: "Product 3", Price: 300},
	//	{ID: uuid.New().String(), Name: "Product 4", Price: 400},
	//}
	//db.Create(&products)

	// read one
	//var product Product
	//db.First(&product)
	//db.First(&product, "name = ?", "Product 1")
	//fmt.Println(product)

	// select all
	//var products []Product
	//db.Limit(2).Offset(2).Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	//where
	//var products []Product
	//db.Where("price > ?", 200).Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	//like
	//var products2 []Product
	//db.Where("name LIKE ?", "%2%").Find(&products2)
	//for _, product := range products2 {
	//	fmt.Println(product)
	//}

	//var p Product
	//db.First(&p, "id = ?", "910319bc-887a-4966-be24-30abeb9b9b8f")
	//p.Name = "New Mouse"
	//
	//db.Save(&p)
	//
	//var p2 Product
	//db.First(&p2, "id = ?", id)
	//fmt.Println(p2.Name)
	//db.Delete(&p)

}
