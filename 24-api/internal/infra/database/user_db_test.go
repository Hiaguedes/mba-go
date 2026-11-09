package database

import (
	"testing"

	"github.com/hiaguedes/go-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("John Doe", "test@test.com", "password")
	userDB := NewUserDB(db)

	err = userDB.Create(user)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail("test@test.com")
	if err != nil {
		t.Fatalf("failed to find user: %v", err)
	}

	assert.Equal(t, user.ID, userFound.ID)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("John Doe", "test@test.com", "password")
	userDB := NewUserDB(db)

	err = userDB.Create(user)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	userFound, err := userDB.FindByEmail("test@test.com")
	if err != nil {
		t.Fatalf("failed to find user: %v", err)
	}

	assert.Equal(t, user.ID, userFound.ID)		
}

func TestFindProductByID(t *testing.T) {
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

	productFound, err := productDB.FindByID(product.ID.String())
	if err != nil {
		t.Fatalf("failed to find product: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
}

func TestUpdateProduct(t *testing.T) {
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

	product.Name = "Updated Product"
	err = productDB.Update(product)
	if err != nil {
		t.Fatalf("failed to update product: %v", err)
	}

	productFound, err := productDB.FindByID(product.ID.String())
	if err != nil {
		t.Fatalf("failed to find product: %v", err)
	}
	assert.NoError(t, err)
	assert.Equal(t, "Updated Product", productFound.Name)
	assert.Equal(t, product.ID, productFound.ID)
}

func TestDeleteProduct(t *testing.T) {
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

	err = productDB.Delete(product.ID.String())
	if err != nil {
		t.Fatalf("failed to delete product: %v", err)
	}

	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
}