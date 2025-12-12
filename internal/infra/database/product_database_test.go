package database

import (
	"strconv"
	"testing"

	"github.com/luis13005/pos-go/internal/model"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	entrada := &model.Product{
		Nome:  "Notebook",
		Preco: decimal.NewFromFloat(2500.50),
	}
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&model.Product{})

	productDB := NewProductDB(db)
	product, err := model.NewProduct(entrada)

	assert.NoError(t, err)

	err = productDB.CreateProduct(product)

	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)

	db.AutoMigrate(&model.Product{})

	entrada := &model.Product{
		Nome:  "NOTEBOOK",
		Preco: decimal.NewFromFloat(2500.50),
	}

	productDB := NewProductDB(db)

	for i := 0; i <= 2; i++ {
		entrada.Nome = entrada.Nome + " " + strconv.Itoa(i)
		product, err := model.NewProduct(entrada)
		assert.NoError(t, err)

		err = productDB.CreateProduct(product)
		assert.NoError(t, err)
	}

	products, err := productDB.FindAll(0, 0, "")

	assert.NoError(t, err)
	assert.NotNil(t, products)
}

func TestFindById(t *testing.T) {
	entrada := &model.Product{
		Nome:  "NOTEBOOK",
		Preco: decimal.NewFromFloat(2500.50),
	}

	db, err := gorm.Open(sqlite.Open("file::momery:"), &gorm.Config{})
	assert.Nil(t, err)

	db.AutoMigrate(&model.Product{})

	productDB := NewProductDB(db)
	product, err := model.NewProduct(entrada)
	assert.NoError(t, err)

	err = productDB.CreateProduct(product)
	assert.NoError(t, err)

	produtoAchado, err := productDB.FindById(product.ID.String())

	assert.NoError(t, err)
	assert.Equal(t, product.ID, produtoAchado.ID)
}

func TestUpdate(t *testing.T) {
	entrada := &model.Product{
		Nome:  "NOTEBOOK",
		Preco: decimal.NewFromFloat(2500.50),
	}
	db, err := gorm.Open(sqlite.Open("file::momery:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&model.Product{})

	productDB := NewProductDB(db)
	product, err := model.NewProduct(entrada)
	assert.NoError(t, err)
	err = productDB.CreateProduct(product)
	assert.NoError(t, err)

	product.Nome = "PC Gamer"
	err = productDB.Update(product)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	entrada := &model.Product{
		Nome:  "NOTEBOOK",
		Preco: decimal.NewFromFloat(2500.50),
	}
	db, err := gorm.Open(sqlite.Open("file::momery:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&model.Product{})
	productDB := NewProductDB(db)
	product, err := model.NewProduct(entrada)
	assert.NoError(t, err)
	err = productDB.CreateProduct(product)
	assert.NoError(t, err)
	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)
}