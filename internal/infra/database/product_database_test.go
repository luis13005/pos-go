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

// func (productDB *ProductDB) FindById(id string) (*model.Product, error) {
// 	var product model.Product

// 	err := productDB.DB.Where("id = ?", id).First(&product).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &product, nil
// }

func TestFindById(t *testing.T) {
	// entrada := &model.Product{
	// 	Nome:  "NOTEBOOK",
	// 	Preco: decimal.NewFromFloat(2500.50),
	// }

	db, err := gorm.Open(sqlite.Open("file::momery:"), &gorm.Config{})
	assert.Nil(t, err)

	db.AutoMigrate(&model.Product{})

}
