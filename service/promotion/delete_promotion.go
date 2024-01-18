package promotion

func (ps *PromotionService) Delete(promotionId string) error {
	clients, err := ps.tokenStorage.GetAll()
	if err != nil {
		return err
	}

	for _, cli := range clients {
		if cli.PromotionId == promotionId {
			return ClientInUsedError
		}
	}

	if err := ps.promotionRepo.Delete(promotionId); err != nil {
		return err
	}

	return nil
}
