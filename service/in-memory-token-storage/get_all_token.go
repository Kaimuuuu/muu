package token

import "kaimuu/model"

func (ts *InMemoryTokenStorage) GetAll() ([]model.Client, error) {
	clients := make([]model.Client, 0)
	for _, c := range ts.store {
		clients = append(clients, *c)
	}

	return clients, nil
}
