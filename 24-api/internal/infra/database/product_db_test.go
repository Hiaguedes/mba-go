package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/hiaguedes/go-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
		db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Product 1", 10.0)
	productDB := NewProductDB(db)

	err = productDB.Create(product)
	if err != nil {
		t.Fatalf("failed to create product: %v", err)
	}

	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&entity.Product{})
	// create a list of  30 products

	productDB := NewProductDB(db)

	for i := 1; i <= 30; i++ {
		product, _ := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)

		err = productDB.Create(product)

		assert.NoError(t, err)
		if err != nil {
			t.Fatalf("failed to create product: %v", err)
		}
	}

	products, err := productDB.FindAll(1, 10, "asc")

	if err != nil {
		t.Fatalf("failed to find products: %v", err)
	}

	assert.Len(t, products, 10)
	assert.NoError(t, err)
	assert.Equal(t, "Product 5", products[4].Name)
	assert.Equal(t, "Product 1", products[0].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	if err != nil {
		t.Fatalf("failed to find products: %v", err)
	}

	assert.Len(t, products, 10)
	assert.NoError(t, err)
	assert.Equal(t, "Product 15", products[4].Name)
}