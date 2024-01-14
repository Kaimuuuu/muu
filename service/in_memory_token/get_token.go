package token

import (
	"fmt"
	"kaimuu/model"
)

func (ts *InMemoryTokenStorage) Get(token string) (*model.Client, error) {
	val, ok := ts.store[token]
	if !ok {
		return &model.Client{}, fmt.Errorf("invalid token: %s", token)
	}

	return val, nil
}
