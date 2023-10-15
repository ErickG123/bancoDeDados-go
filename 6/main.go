package main

import (
	"gorm.io/driver/mysql" // Driver do MySQL
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	// Trabalhando com Lock
	// Iniciando uma Transaction
	tx := db.Begin()
	var c Category
	// Bloquando a linha para realizar um UPDATE
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}

	// Fazendo um UPDATE e dando COMMIT na TRANSACTION
	c.Name = "Eletronicos"
	tx.Debug().Save(&c)
	tx.Commit()
}
