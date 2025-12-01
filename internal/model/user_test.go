package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	entrada := User{
		Nome:  "nome teste",
		Email: "email@teste.com",
		Senha: "admin",
	}

	user, err := NewUser(entrada)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Senha)
	assert.Equal(t, "nome teste", user.Nome)
	assert.Equal(t, "admin", user.Email)
	validou, _ := user.ValidaSenha("admin")
	assert.True(t, validou)
}
