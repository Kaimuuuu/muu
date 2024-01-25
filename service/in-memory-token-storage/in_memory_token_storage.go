package token

import (
	"kaimuu/model"

	"github.com/cockroachdb/errors"
)

var (
	InvalidTokenError = errors.New("invalid token")
)

type InMemoryTokenStorage struct {
	store map[string]*model.Client
}

func NewInMemoryTokenStorage() *InMemoryTokenStorage {
	return &InMemoryTokenStorage{
		store: make(map[string]*model.Client),
	}
}
