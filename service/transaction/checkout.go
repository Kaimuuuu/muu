package transaction

func (ts *TransactionService) Checkout(token string) error {
	c, err := ts.tokenRepo.Get(token)
	if err != nil {
		return err
	}

	t, err := ts.toTransaction(c)
	if err != nil {
		return err
	}

	if err := ts.transactionRepo.Insert(t); err != nil {
		return err
	}

	if err := ts.tokenRepo.Delete(token); err != nil {
		return err
	}

	return nil
}
