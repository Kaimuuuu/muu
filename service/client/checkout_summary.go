package client

func (cs *ClientService) CheckoutSummary(token string) (*TransactionObject, error) {
	cli, err := cs.tokenStorage.Get(token)

	if err != nil {
		return &TransactionObject{}, err
	}

	trans, err := cs.toTransactionObject(cli)

	if err != nil {
		return &TransactionObject{}, err
	}

	return trans, err
}
