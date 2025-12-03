package database

import (
	"github.com/luis13005/pos-go/internal/model"
)

type UserInterface interface {
	CreateUser(user *model.User) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}
