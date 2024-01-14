package promotion

func (ps *PromotionService) Delete(promotionId string) error {
	if err := ps.promotionRepo.Delete(promotionId); err != nil {
		return err
	}

	return nil
}
