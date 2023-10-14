package main

// Utilizando o pacote da Google para gerar UUID
import (
	"database/sql"
	"fmt"

	// Faz com que o programa compile o código sem excluir esse pacote
	// O "_" faz com que a linha não suma
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

// Criando uma Struct para a tabela
type Product struct {
	ID    string
	Name  string
	Price float64
}

// Criando uma função para criar um novo Produto
func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	// Se conectando com o Banco de Dados
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}

	// Fechando a conexão com o Banco de Dados
	defer db.Close()

	// Criando um novo Produto
	product := NewProduct("Celular", 2459.90)
	// Inserindo o Produto
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	// Alterando o preço do Produto
	product.Price = 100.0
	// Fazendo o UPDATE do Produto
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	// Selecionando um Produto
	p, err := selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Produto: %v, custa R$%.2f \n", p.Name, p.Price)

	// Selecionando todos os Produtos
	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}

	// Percorrendo o Slice para mostrar cada Produto
	for _, p := range products {
		fmt.Printf("Produto: %v, custa R$%.2f \n", p.Name, p.Price)
	}

	// Deletando o Produto
	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
}

// Criando uma função para Inserir um novo dado
func insertProduct(db *sql.DB, product *Product) error {
	// Prepar Statement, protege o processo de inserção de dados
	stmt, err := db.Prepare("INSERT INTO products(id, name, price) values(?, ?,?)")
	if err != nil {
		panic(err)
	}

	// Finalizar o Statement no final
	defer stmt.Close()

	// Executando o INSERT
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		panic(nil)
	}

	fmt.Printf("Produto %v cadastrado com Sucesso. \n", product.Name)

	return nil
}

// Criando uma função para fazer um UPDATE
func updateProduct(db *sql.DB, product *Product) error {
	// Preparando o UPDATE
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	// Executando o UPDATE
	// Eu passo os parâmetros na ordem do Prepare
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Produto %v atualizado com Sucesso. \n", product.Name)

	return nil
}

// Criando uma função para Selecionar dados de um Produto
func selectProduct(db *sql.DB, id string) (*Product, error) {
	// Preparando o SELECT
	stmt, err := db.Prepare("SELECT * FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// Criando uma variável vazia do tipo Product
	var p Product
	// QueryRow busca apenas uma linha
	// O Scan vai analizar cada coluna da tabela e atribuir cada valor
	// a uma campo da nossa Struct de Produto
	// Fazendo isso, eu estou mudando o valor de "p"
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}

	// Existe a possibilidade de eu passar um Context no QueryRow
	// para isso eu utilizo o QueryRowContext

	// Retornando o Endereço de "p"
	return &p, nil
}

// Função para Selecionar todos os Produtos
func selectAllProducts(db *sql.DB) ([]Product, error) {
	// Neste caso, não preciso fazer o Prepare, já que não
	// vou passar parametros
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	// Fechando a conexão com o banco de dados
	defer rows.Close()

	// Guardando um Slice de Produtos
	var products []Product

	// Percorre linha por linha
	for rows.Next() {
		// Salvando cada Produto no p
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			panic(err)
		}

		// Salvando o valor de P no Slice de Produtos
		products = append(products, p)
	}

	return products, nil
}

// Função para Deleter um Produto
func deleteProduct(db *sql.DB, id string) error {
	// Preparando o DELETE
	stmt, err := db.Prepare("DELETE FROM produts WHERE id = ?")
	if err != nil {
		panic(err)
	}

	// Fechando a Conexão com o Banco
	defer stmt.Close()

	// Executando o DELETE
	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}

	return nil
}
