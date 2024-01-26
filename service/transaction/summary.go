package transaction

import "kaimuu/model"

func (ts *TransactionService) Summary(token string) (*model.Transaction, error) {
	c, err := ts.tokenRepo.Get(token)
	if err != nil {
		return &model.Transaction{}, err
	}

	t, err := ts.toTransaction(c)
	if err != nil {
		return &model.Transaction{}, err

	}

	return t, err
}
