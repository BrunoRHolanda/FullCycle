package application_test

import (
	"github.com/BrunoRHolanda/FullCycle/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0

	err = product.Enable()

	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10

	err = product.Disable()

	require.Equal(t, "the price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{
		ID:     uuid.NewV4().String(),
		Name:   "Hello",
		Status: application.DISABLED,
		Price:  0,
	}

	_, err := product.IsValid()

	require.Nil(t, err)

	product.Status = "INVALID"

	_, err = product.IsValid()

	require.Equal(t, "the status must be ENABLED or DISABLED", err.Error())
}
