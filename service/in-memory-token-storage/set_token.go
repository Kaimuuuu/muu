package token

import (
	"kaimuu/model"

	"github.com/cockroachdb/errors"
)

func (ts *InMemoryTokenStorage) Set(token string, c *model.Client) error {
	_, ok := ts.store[token]
	if ok {
		return errors.Newf("token '%s' already exist", token)
	}

	ts.store[token] = c
	return nil
}
