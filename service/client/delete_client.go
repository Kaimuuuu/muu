package client

func (cs *ClientService) Delete(token string) error {
	if err := cs.tokenStorage.Remove(token); err != nil {
		return err
	}

	return nil
}
