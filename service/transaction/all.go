package transaction

import "github.com/Kaimuuuu/muu/model"

func (ts *TransactionService) All() ([]model.Transaction, error) {
	transactions, err := ts.transactionRepo.All()
	if err != nil {
		return []model.Transaction{}, err
	}

	return transactions, nil
}
