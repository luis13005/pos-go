package database

import (
	"github.com/luis13005/pos-go/internal/model"
	"gorm.io/gorm"
)

type ProductDB struct {
	*gorm.DB
}

func NewProductDB(db *gorm.DB) *ProductDB {
	return &ProductDB{db}
}

func (productDb *ProductDB) CreateProduct(p *model.Product) error {

	if err := productDb.DB.Create(p).Error; err != nil {
		return err
	}

	return nil
}

func (productDb *ProductDB) FindAll(page, limit int, sort string) ([]model.Product, error) {
	var products []model.Product
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = productDb.DB.Limit(limit).Offset((page - 1) * limit).Order("Nome " + sort).Find(&products).Error
	} else {
		err = productDb.DB.Order("Nome " + sort).Find(&products).Error
	}

	return products, err
}

// FindAll(page, limit int, sort string) ([]model.Product, error)
// FindById(id string) (*model.Product, error)
// Update(product *model.Product) error
// Delete(id string) error
