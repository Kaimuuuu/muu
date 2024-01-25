package token

import (
	"kaimuu/model"
)

func (ts *InMemoryTokenStorage) Get(token string) (*model.Client, error) {
	val, ok := ts.store[token]
	if !ok {
		return &model.Client{}, InvalidTokenError
	}

	return val, nil
}
