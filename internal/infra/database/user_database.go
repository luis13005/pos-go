package database

import (
	"github.com/luis13005/pos-go/internal/model"
	"gorm.io/gorm"
)

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{db: db}
}

func (u *UserDB) CreateUser(user *model.User) (*model.User, error) {
	err := u.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserDB) FindByEmail(email string) (*model.User, error) {
	var usuario model.User
	if err := u.db.Where("email = ?", email).First(&usuario).Error; err != nil {
		return nil, err
	}

	return &usuario, nil
}
