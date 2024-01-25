package promotion

func (ps *PromotionService) GetWeight(promotionId string) (int, error) {
	p, err := ps.promotionRepo.GetById(promotionId)
	if err != nil {
		return 0, err
	}

	return p.Weight, nil
}
