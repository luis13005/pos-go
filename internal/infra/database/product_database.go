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

func (productDB *ProductDB) CreateProduct(p *model.Product) error {

	if err := productDB.DB.Create(p).Error; err != nil {
		return err
	}

	return nil
}

func (productDB *ProductDB) FindAll(page, limit int, sort string) ([]model.Product, error) {
	var products []model.Product
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = productDB.DB.Limit(limit).Offset((page - 1) * limit).Order("Nome " + sort).Find(&products).Error
	} else {
		err = productDB.DB.Order("Nome " + sort).Find(&products).Error
	}

	return products, err
}

func (productDB *ProductDB) FindById(id string) (*model.Product, error) {
	var product model.Product

	err := productDB.DB.Where("id = ?", id).First(&product).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (productDB *ProductDB) Update(product *model.Product) error {
	return productDB.DB.Save(product).Error
}

func (productDB *ProductDB) Delete(id string) error {
	product, err := productDB.FindById(id)
	if err != nil {
		return err
	}

	return productDB.DB.Delete(product).Error
}
