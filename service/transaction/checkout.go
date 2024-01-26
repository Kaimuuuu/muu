package transaction

func (ts *TransactionService) Checkout(token string) error {
	c, err := ts.tokenServ.Get(token)
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

	if err := ts.tokenServ.Delete(token); err != nil {
		return err
	}

	return nil
}
