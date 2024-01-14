package token

import (
	"fmt"
	"kaimuu/model"
)

func (ts *InMemoryTokenStorage) Set(token string, cli *model.Client) error {
	_, ok := ts.store[token]
	if ok {
		return fmt.Errorf("token already exist: %s", token)
	}

	ts.store[token] = cli
	return nil
}
