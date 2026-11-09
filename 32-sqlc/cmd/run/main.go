package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hiaguedes/sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	queries := db.New(conn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Programming",
	// 	Description: sql.NullString{String: "Description", Valid: true},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Original Category: ", categories)

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		Name: "BackEnd With Golang",
		Description: sql.NullString{String: "Description for BackEnd With Golang", Valid: true},
		ID: "377e7ad7-513c-477d-944c-466ec665477e",
	})

	if err != nil {
		panic(err)
	}

	updatedCategories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated Category: ", updatedCategories)
}