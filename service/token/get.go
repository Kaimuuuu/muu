package token

import "github.com/Kaimuuuu/muu/model"

func (ts *TokenService) Get(token string) (*model.Client, error) {
	c, err := ts.tokenRepo.Get(token)
	if err != nil {
		return &model.Client{}, err
	}

	return c, nil
}
