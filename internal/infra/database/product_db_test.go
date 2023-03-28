package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/hpaes/api-project-golang/internal/entity"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGivenValidParamsShouldInsertProductInDb(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&entity.Product{})

	productRepo := NewProductRepository(db)
	product, err := entity.NewProduct("Product 1", "Description 1", 10.0)
	assert.NoError(t, err)

	err = productRepo.Create(product)
	assert.NoError(t, err)

	var productFound entity.Product
	err = db.Find(&productFound, "id = ?", product.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Description, productFound.Description)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestGivenValidIdShouldFindProductFindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&entity.Product{})

	productRepo := NewProductRepository(db)
	product, err := entity.NewProduct("Product 1", "Description 1", 10.0)
	assert.NoError(t, err)

	err = productRepo.Create(product)
	assert.NoError(t, err)

	productFound, err := productRepo.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Description, productFound.Description)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestGivenValidParamsShouldUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&entity.Product{})

	productRepo := NewProductRepository(db)
	product, err := entity.NewProduct("Product 1", "Description 1", 10.0)
	assert.NoError(t, err)

	err = productRepo.Create(product)
	assert.NoError(t, err)

	productFound, err := productRepo.FindById(product.ID.String())
	assert.NoError(t, err)

	productToUpdate := &entity.Product{
		ID:          productFound.ID,
		Name:        "Product 2",
		Description: "Description 2",
		Price:       20.0,
	}

	err = productRepo.Update(productToUpdate)
	assert.NoError(t, err)

	productFound, err = productRepo.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, productFound.ID, productToUpdate.ID)
	assert.Equal(t, productFound.Name, productToUpdate.Name)
	assert.Equal(t, productFound.Description, productToUpdate.Description)
	assert.Equal(t, productFound.Price, productToUpdate.Price)
}

func TestGivenValidIdShouldDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&entity.Product{})

	productRepo := NewProductRepository(db)
	product, err := entity.NewProduct("Product 1", "Description 1", 10.0)
	assert.NoError(t, err)

	err = productRepo.Create(product)
	assert.NoError(t, err)

	err = productRepo.Delete(product.ID.String())
	assert.NoError(t, err)

	productFound, err := productRepo.FindById(product.ID.String())
	assert.Error(t, err)
	assert.Nil(t, productFound)
}

func TestGivenNoParametersShouldFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&entity.Product{})

	productRepo := NewProductRepository(db)

	for i := 0; i < 10; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), fmt.Sprintf("Description %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		err = productRepo.Create(product)
		assert.NoError(t, err)
	}

	products, err := productRepo.FindAll(0, 0, "")
	assert.NoError(t, err)
	assert.Equal(t, 10, len(products))
}

func TestGivenPageLimitAndSortShouldFindOnlyTheLimitOfProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&entity.Product{})

	productRepo := NewProductRepository(db)

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), fmt.Sprintf("Description %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		err = productRepo.Create(product)
		assert.NoError(t, err)
	}

	products, err := productRepo.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productRepo.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productRepo.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}
