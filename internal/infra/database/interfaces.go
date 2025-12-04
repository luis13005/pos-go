package database

import (
	"github.com/luis13005/pos-go/internal/model"
)

type UserInterface interface {
	CreateUser(user *model.User) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

type ProductInterface interface {
	CreateProduct(product *model.Product) error
	FindAll(page, limit int, sort string) ([]model.Product, error)
	FindById(id string) (*model.Product, error)
	Update(product *model.Product) error
	Delete(id string) error
}
