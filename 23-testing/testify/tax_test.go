package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	// arrange
	amount := 500.0
	// act
	tax, err := CalculateTax(amount)
	// assert
	assert.Equal(t, 5.0, tax)
	assert.NoError(t, err)
	assert.NotEqual(t, 10.0, tax)

	// arrange
	amount = 0
	// act
	tax, err = CalculateTax(amount)

	// assert
	assert.Equal(t, 0.0, tax)
	assert.Error(t, err)
	assert.EqualError(t, err, "amount must be greater than zero")
}

func TestCalculateTaxAndSave(t *testing.T) {
	// arrange
	amount := 500.0
	repoMock := &TaxRepositoryMock{}
	repoMock.On("SaveTax", 5.0).Return(nil)

	// act
	err := CalculateTaxAndSave(amount, repoMock)
	// assert
	assert.Nil(t, err)

	// act with an amount that will cause an error
	err = CalculateTaxAndSave(0.0, repoMock)

	// assert
	assert.Error(t, err, "amount must be greater than zero")

	repoMock.AssertExpectations(t)
	repoMock.AssertNumberOfCalls(t, "SaveTax", 1)
}