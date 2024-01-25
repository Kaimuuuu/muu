package client

func (cs *ClientService) CheckoutSummary(token string) (*TransactionObject, error) {
	c, err := cs.tokenStorage.Get(token)
	if err != nil {
		return &TransactionObject{}, err
	}

	t, err := cs.toTransactionObject(c)
	if err != nil {
		return &TransactionObject{}, err

	}

	return t, err
}
