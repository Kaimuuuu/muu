package token

import "kaimuu/model"

func (ts *TokenService) Get(token string) (*model.Client, error) {
	c, err := ts.tokenRepo.Get(token)
	if err != nil {
		return &model.Client{}, err
	}

	return c, nil
}
