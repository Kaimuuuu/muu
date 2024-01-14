package token

import "fmt"

func (ts *InMemoryTokenStorage) Remove(token string) error {
	_, ok := ts.store[token]
	if !ok {
		return fmt.Errorf("invalid token: %s", token)
	}

	delete(ts.store, token)
	return nil
}
