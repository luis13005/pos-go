package model

import (
	"errors"
	"time"

	"github.com/luis13005/pos-go/pkg/entidade"
	"github.com/shopspring/decimal"
)

var (
	ErrIdIsRequired    = errors.New("id is required")
	ErrInvalidId       = errors.New("invalid id")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID        entidade.ID     `json:"id"`
	Nome      string          `json:"nome"`
	Preco     decimal.Decimal `json:"preco"`
	CreatedAt string          `json:"created_at"`
}

func NewProduct(p *Product) (*Product, error) {
	produto := &Product{
		ID:        entidade.NewId(),
		Nome:      p.Nome,
		Preco:     p.Preco,
		CreatedAt: time.Now().String(),
	}

	err := produto.Validate()
	if err != nil {
		return nil, err
	}

	return produto, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIdIsRequired
	}

	if _, err := entidade.Parse(p.ID.String()); err != nil {
		return ErrInvalidId
	}

	if p.Nome == "" {
		return ErrNameIsRequired
	}

	if p.Preco.IsZero() {
		return ErrPriceIsRequired
	}

	if p.Preco.IsNegative() {
		return ErrInvalidPrice
	}

	return nil
}
