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
	user, err := model.NewUser(entrada)

	assert.NoError(t, err)
	usuario, err := userDB.CreateUser(user)

	assert.NoError(t, err)
	assert.NotEmpty(t, usuario.ID)
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

	assert.NoError(t, err)
	assert.NotNil(t, userByEmail)

	userByEmail, err = userDb.FindByEmail("tes@tee.com")

	assert.NotNil(t, err)
	assert.Nil(t, userByEmail)
}
