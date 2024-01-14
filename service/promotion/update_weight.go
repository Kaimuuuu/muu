package promotion

func (ps *PromotionService) UpdateWeight(promotionId string, weight int) error {
	promo, err := ps.promotionRepo.GetById(promotionId)

	if err != nil {
		return err
	}

	promo.Weight = weight

	if err := ps.promotionRepo.Update(promotionId, promo); err != nil {
		return err
	}

	return nil
}
