package promotion

func (ps *PromotionService) UpdatePromotion(promotionId string, req UpdatePromotionRequest) error {
	clients, err := ps.tokenStorage.GetAll()
	if err != nil {
		return err
	}

	for _, cli := range clients {
		if cli.PromotionId == promotionId {
			return ClientInUsedError
		}
	}

	promo, err := ps.promotionRepo.GetById(promotionId)

	if err != nil {
		return err
	}

	promo.Name = req.Name
	promo.Weight = req.Weight
	promo.Description = req.Description
	promo.Price = req.Price
	promo.Duration = req.Duration
	promo.PromotionMenuItems = req.PromotionMenuItems
	promo.ImagePath = req.ImagePath

	if err := ps.promotionRepo.Update(promotionId, promo); err != nil {
		return err
	}

	return nil
}
