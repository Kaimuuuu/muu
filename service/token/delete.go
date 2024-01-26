package token

func (ts *TokenService) Delete(token string) error {
	if err := ts.tokenRepo.Delete(token); err != nil {
		return err
	}

	return nil
}
