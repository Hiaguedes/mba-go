//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/hiaguedes/DI/product"
)

// se precisar de um product repostiory txt eu consigo reaproveitar o setRepositoryDependency e criar um novo setRepositoryDependencyTxt, e assim consigo injetar a dependencia do repository txt no useCase
var setRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)


func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepositoryDependency,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
