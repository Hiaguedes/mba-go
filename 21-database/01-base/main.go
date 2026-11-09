package main

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql" // _ is a blank identifier
	"github.com/google/uuid"
)

type Product struct {
	Id    string
	Name  string
	Price float64
}

var (
	dbInstance *sql.DB
	once       sync.Once
)

func NewProduct(name string, price float64) *Product {
	return &Product{
		Id: uuid.New().String(),
		Name: name,
		Price: price,
	}
}

func main() {
	db := GetDB()

	// product := &Product{
	// 	Id: "e441ed88-ba42-48e0-af99-358fd0d6e366",
	// 	Name: "Cereal",
	// 	Price: 139.99,
	// }
	// p, err := selectOneProduct(db, product.Id)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Product %+v", p)

	// products, err := selectAllProducts(db)
	// if err != nil {
	// 	panic(err)
	// }

	// for _,p := range products {
	// 	fmt.Printf("Product: %+v", p)
	// }

	err := deleteProduct(db, "bce0c50a-d1ba-4404-8f9c-2b16540d8f03")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}

func GetDB() *sql.DB {
	once.Do(func() {
		var err error
		// Open the database connection
		dbInstance, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
		if err != nil {
			log.Fatalf("Error opening database: %v", err)
		}

		// Verify the connection is actually working
		err = dbInstance.Ping()
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
	})

	return dbInstance
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Id, product.Name, product.Price)

		if err != nil {
		return err
	}

	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.Id)

		if err != nil {
		return err
	}

	return nil
}

func selectOneProduct(db *sql.DB, id string) (*Product, error){
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var product Product
	err = stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price)

	if err != nil {
		return nil, err
	}

	return &product, nil

}

func selectAllProducts(db *sql.DB) ([]Product,error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []Product
	for rows.Next(){
		var p Product
		err = rows.Scan(&p.Id, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}