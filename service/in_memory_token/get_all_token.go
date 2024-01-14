package token

import "kaimuu/model"

func (ts *InMemoryTokenStorage) GetAll() ([]model.Client, error) {
	clients := make([]model.Client, 0)
	for _, cli := range ts.store {
		clients = append(clients, *cli)
	}

	return clients, nil
}
