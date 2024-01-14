package client

import "kaimuu/model"

func (cs *ClientService) GetClients() ([]model.Client, error) {
	clients, err := cs.tokenStorage.GetAll()

	if err != nil {
		return []model.Client{}, err
	}

	return clients, nil
}
