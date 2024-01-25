package client

func (cs *ClientService) Checkout(token string) error {
	c, err := cs.tokenStorage.Get(token)
	if err != nil {
		return err
	}

	t, err := cs.toTransactionObject(c)
	if err != nil {
		return err
	}

	if err := cs.transactionRepo.Insert(t); err != nil {
		return err
	}

	if err := cs.tokenStorage.Remove(token); err != nil {
		return err
	}

	return nil
}
