package main

import (
	"fmt"

	"gorm.io/driver/mysql" // Driver do MySQL
	"gorm.io/gorm"
)

// Criando um Relacionamento N:N
type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
	gorm.Model
}

func main() {
	// Conexão com o Banco de Dados
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	// Trabalhando com Relacionamentos
	// many_to_many => Relacionamento N:N
	category := Category{Name: "Cozinha"}
	db.Create(&category)

	category2 := Category{Name: "Eletronicos"}
	db.Create(&category2)

	db.Create(&Product{
		Name:       "Mesa",
		Price:      1000.00,
		Categories: []Category{category, category2},
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name)
		for _, product := range category.Products {
			println("- ", product.Name)
		}
	}
}
