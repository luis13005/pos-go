package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestNewUser(t *testing.T) {
	entrada := &User{
		Nome:  "nome teste",
		Email: "email@teste.com",
	}

	user, err := NewUser(entrada)
	assert.Error(t, err, "Esperado erro ao tentar gerar hash de senha vazia")
	assert.Nil(t, user, "User deve ser nil quando há erro")

	entrada.Senha = "admin"
	user, err = NewUser(entrada)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)

	assert.Equal(t, "nome teste", user.Nome)
	assert.Equal(t, "email@teste.com", user.Email)

	//SENHA
	assert.NotEmpty(t, user.Senha)
	validou, _ := user.ValidaSenha("admin")
	assert.True(t, validou)
	validou, err = user.ValidaSenha("123123")
	assert.False(t, validou)
	assert.ErrorIs(t, err, bcrypt.ErrMismatchedHashAndPassword, "Esperado erro de hash/senha incorreta")
}
