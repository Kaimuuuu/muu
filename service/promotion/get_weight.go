package promotion

func (ps *PromotionService) GetWeight(promotionId string) (int, error) {
	promo, err := ps.promotionRepo.GetById(promotionId)
	if err != nil {
		return 0, err
	}

	return promo.Weight, nil
}
