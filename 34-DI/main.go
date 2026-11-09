package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// encadeamento de dependências
	// repository := product.NewProductRepository(db)
	// useCase := product.NewProductUseCase(repository)

	// no final voce pode usar o wire para injetar as dependências e voce so precisa do useCase, sem se preocupar com o repository
	useCase := NewUseCase(db)
	
	product, err := useCase.GetProductByID(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(product.Name)

}