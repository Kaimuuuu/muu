package promotion

func (ps *PromotionService) UpdatePromotion(promotionId string, req UpdatePromotionRequest) error {
	promo, err := ps.promotionRepo.GetById(promotionId)

	if err != nil {
		return err
	}

	promo.Name = req.Name
	promo.Description = req.Description
	promo.Price = req.Price
	promo.Duration = req.Duration
	promo.PromotionMenuItems = req.PromotionMenuItems

	if err := ps.promotionRepo.Update(promotionId, promo); err != nil {
		return err
	}

	return nil
}
