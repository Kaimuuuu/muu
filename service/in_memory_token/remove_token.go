package token

func (ts *InMemoryTokenStorage) Remove(token string) error {
	_, ok := ts.store[token]
	if !ok {
		return InvalidTokenError
	}

	delete(ts.store, token)
	return nil
}
