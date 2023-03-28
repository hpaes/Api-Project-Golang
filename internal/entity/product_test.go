package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenValidParamShouldCreateProduct(t *testing.T) {
	product, err := NewProduct("Product 1", "Product 1 Description", 100)
	assert.Nil(t, err)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, "Product 1 Description", product.Description)
	assert.Equal(t, 100.00, product.Price)

	err = product.Validate()
	assert.NoError(t, err)
}

func TestGivenInvalidNameShouldNotCreateProduct(t *testing.T) {
	product, err := NewProduct("", "", 100)
	assert.Error(t, err, "name is required")
	assert.Nil(t, product)
}

func TestGivenInvalidPriceShouldNotCreateProduct(t *testing.T) {
	product, err := NewProduct("Product 1", "Product 1 Description", -1)
	assert.Error(t, err, "invalid price")
	assert.Nil(t, product)
}

// func TestGivenInvalidQuantityShouldNotCreateProduct(t *testing.T) {
// 	product, err := NewProduct("Product 1", "Product 1 Description", 100)
// 	assert.Error(t, err, "quantity is required")
// 	assert.Nil(t, product)
// }
