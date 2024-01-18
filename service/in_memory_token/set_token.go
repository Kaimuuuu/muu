package token

import (
	"kaimuu/model"

	"github.com/cockroachdb/errors"
)

func (ts *InMemoryTokenStorage) Set(token string, cli *model.Client) error {
	_, ok := ts.store[token]
	if ok {
		return errors.Newf("token '%s' already exist", token)
	}

	ts.store[token] = cli
	return nil
}
