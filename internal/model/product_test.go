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

	produto, err := NewProduct(*entrada)
	assert.Nil(t, err)
	assert.NotNil(t, produto)
	assert.Equal(t, "notebook", produto.Nome)
	assert.Equal(t, decimal.NewFromFloat(2500.50), produto.Preco)
}
