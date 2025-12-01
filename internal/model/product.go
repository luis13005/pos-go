package model

import "github.com/luis13005/pos-go/pkg/entidade"

type Product struct {
	ID   entidade.ID `json:"id"`
	Nome string      `json:"nome"`
}
