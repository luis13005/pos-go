package database

import (
	"testing"

	"github.com/luis13005/pos-go/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	entrada := &model.User{
		Nome:  "Luis Fernando",
		Email: "teste@teste.com",
		Senha: "admin",
	}

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&model.User{})
	userDB := NewUserDB(db)

	usuario, err := userDB.CreateUser(entrada)

	assert.Nil(t, err)
	assert.NotNil(t, usuario)
}

func TestFindByEmail(t *testing.T) {
	entrada := &model.User{
		Nome:  "Luis Fernando",
		Email: "teste@teste.com",
		Senha: "admin",
	}

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&model.User{})

	userDb := NewUserDB(db)

	_, err = userDb.CreateUser(entrada)
	assert.Nil(t, err)

	userByEmail, err := userDb.FindByEmail("teste@teste.com")

	assert.Nil(t, err)
	assert.NotNil(t, userByEmail)

	userByEmail, err = userDb.FindByEmail("tes@tee.com")

	assert.Nil(t, err)
	assert.NotNil(t, userByEmail)
}
