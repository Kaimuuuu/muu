package token

import "github.com/Kaimuuuu/muu/model"

func (ts *TokenService) All() ([]model.Client, error) {
	clients, err := ts.tokenRepo.All()
	if err != nil {
		return []model.Client{}, err
	}

	return clients, nil
}
