package token

import "kaimuu/model"

type InMemoryTokenStorage struct {
	store map[string]*model.Client
}

func NewInMemoryTokenStorage() *InMemoryTokenStorage {
	return &InMemoryTokenStorage{
		store: make(map[string]*model.Client),
	}
}
