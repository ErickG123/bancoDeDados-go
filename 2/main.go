package main

import (
	"fmt"

	"gorm.io/driver/mysql" // Driver do MySQL
	"gorm.io/gorm"
)

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	gorm.Model // Gera create_at, updated_at, deleted_at
}

func main() {
	// Utilizando ORM (GORM)
	// Instalação da ORM: go get -u gorm.io/gorm

	// Conexão com o Banco de Dados
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Fazendo isso eu consigo rodar uma Migrate que cria a
	// tabela de Produtos
	db.AutoMigrate(&Product{})

	// Criando um Produto
	db.Create(&Product{
		Name:  "Mouse",
		Price: 369.90,
	})

	// Criando vários Produtos
	products := []Product{
		{Name: "Teclado", Price: 460.00},
		{Name: "Monitor", Price: 450.00},
		{Name: "Fone", Price: 300.00},
	}
	db.Create(&products)

	// Selecionando o Primeiro Produto
	var product Product
	// Busca o Primeiro Produto com ID = 1
	db.First(&product, 1)
	fmt.Println(product)

	// Busca o Primeiro Produto com Name = "mouse"
	db.First(&product, "name = ?", "Mouse")
	fmt.Println(product)

	// Selecionando Todos os Products
	var prodcuts []Product
	db.Find(&products)
	for _, product := range prodcuts {
		fmt.Println(product)
	}

	// Limitar o Find a trazer 2 registros
	var prodcuts2 []Product
	db.Limit(2).Find(&products)
	for _, product := range prodcuts2 {
		fmt.Println(product)
	}

	// Offset é para paginção
	var prodcuts3 []Product
	db.Limit(2).Offset(2).Find(&products)
	for _, product := range prodcuts3 {
		fmt.Println(product)
	}

	// SELECT com WHERE
	var products4 []Product
	db.Where("price > ?", 300).Find(&products4)
	for _, product := range products4 {
		fmt.Println(product)
	}

	// SELECT com LIKE
	var products5 []Product
	db.Where("name LIKE > ?", "%book%").Find(&products5)
	for _, product := range products5 {
		fmt.Println(product)
	}

	// UPDATE
	var p Product
	db.First(&p, 1)
	p.Name = "Microfone"
	db.Save(&p)

	var p2 Product
	db.First(&p, 1)
	fmt.Println(p2.Name)

	// DELETE
	db.Delete(&p2)

	// created_at
	db.Create(&Product{
		Name:  "Mouse",
		Price: 369.90,
	})

	// updated_at
	var p3 Product
	db.First(&p3, 1)
	p.Name = "New Mouse"
	db.Save(&p)

	// SOFT DELETE => deleted_at
	// O Registro não é Excluido do Banco de Dados
	var p4 Product
	db.First(&p4, 1)
	db.Delete(&p4)
}
