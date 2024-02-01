package promotion

func (ps *PromotionService) GetMenu(promotionId string) ([]PromotionMenuItemResponse, error) {
	p, err := ps.promotionRepo.GetById(promotionId)
	if err != nil {
		return []PromotionMenuItemResponse{}, err
	}

	promotionMenuItems := make([]PromotionMenuItemResponse, len(p.PromotionMenuItems))
	for i, pmi := range p.PromotionMenuItems {
		m, err := ps.menuRepo.GetById(pmi.MenuItemId)
		if err != nil {
			return []PromotionMenuItemResponse{}, err
		}

		promotionMenuItems[i].Type = pmi.Type
		promotionMenuItems[i].MenuItem = *m
		promotionMenuItems[i].Limit = pmi.Limit
	}

	return promotionMenuItems, nil
}
