package app_test

import (
    uuid "github.com/satori/go.uuid"
    "github.com/stretchr/testify/require"
    "github.com/williamtrevisan/go-arch-hexagonal/app"
    "testing"
)

func TestProduct_Enable(t *testing.T) {
    product := app.Product{}
    product.Name = "Product"
    product.Price = 10
    product.Status = app.DISABLED

    err := product.Enable()
    require.Nil(t, err)

    product.Price = 0
    err = product.Enable()
    require.Equal(t, "The price must be greater than zero to enable the product.", err.Error())
}

func TestProduct_Disable(t *testing.T) {
    product := app.Product{}
    product.Name = "Product"
    product.Price = 0
    product.Status = app.ENABLED

    err := product.Disable()
    require.Nil(t, err)

    product.Price = 10
    err = product.Disable()
    require.Equal(t, "The price must be zero in order to have the product disabled.", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
    product := app.Product{}
    product.ID = uuid.NewV4().String()
    product.Name = "Product"
    product.Price = 10
    product.Status = app.DISABLED

    _, err := product.IsValid()
    require.Nil(t, err)

    product.Status = "INVALID"
    _, err = product.IsValid()
    require.Equal(t, "The status must be enabled or disabled.", err.Error())

    product.Status = app.ENABLED
    _, err = product.IsValid()
    require.Nil(t, err)

    product.Price = -10
    _, err = product.IsValid()
    require.Equal(t, "The price must be greater or equal zero.", err.Error())
}
