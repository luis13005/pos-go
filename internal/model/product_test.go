package model

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	entrada := &Product{
		Nome:  "notebook",
		Preco: decimal.NewFromFloat(2500.50),
	}

	produto, err := NewProduct(entrada)
	assert.Nil(t, err)
	assert.NotNil(t, produto)
	assert.Equal(t, "notebook", produto.Nome)
	assert.Equal(t, decimal.NewFromFloat(2500.50), produto.Preco)
}

func TestNameIsRequired(t *testing.T) {
	entrada := &Product{
		Preco: decimal.NewFromInt(10),
	}

	produto, err := NewProduct(entrada)
	assert.Nil(t, produto)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestPriceIsRequired(t *testing.T) {
	entrada := &Product{
		Nome: "Teste Nome",
	}

	produto, err := NewProduct(entrada)
	assert.Nil(t, produto)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestInvalidPrice(t *testing.T) {
	entrada := &Product{
		Nome:  "Teste Nome",
		Preco: decimal.NewFromFloat(-15.50),
	}
	produto, err := NewProduct(entrada)
	assert.Nil(t, produto)
	assert.Equal(t, ErrInvalidPrice, err)

}
