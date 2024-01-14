package client

func (cs *ClientService) Checkout(token string) error {
	cli, err := cs.tokenStorage.Get(token)

	if err != nil {
		return err
	}

	trans, err := cs.toTransactionObject(cli)

	if err != nil {
		return err
	}

	if err := cs.transactionRepo.Insert(trans); err != nil {
		return err
	}

	if err := cs.tokenStorage.Remove(token); err != nil {
		return err
	}

	return nil
}
