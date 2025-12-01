package model

import (
	"github.com/luis13005/pos-go/pkg/entidade"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID    entidade.ID `json:"id"`
	Nome  string      `json:"nome"`
	Email string      `json:"email"`
	Senha string      `json:"senha"`
}

func NewUser(u User) (*User, error) {
	senhaHash, err := bcrypt.GenerateFromPassword([]byte(u.Senha), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		entidade.NewId(),
		u.Nome,
		u.Email,
		string(senhaHash),
	}, nil
}

func (u *User) ValidaSenha(senha string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Senha), []byte(senha))
	if err != nil {
		return false, err
	}

	return true, nil
}
