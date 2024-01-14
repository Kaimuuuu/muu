package promotion

func (ps *PromotionService) GetPromotionMenu(promotionId string) ([]PromotionMenuItemResponse, error) {
	prom, err := ps.promotionRepo.GetById(promotionId)

	if err != nil {
		return []PromotionMenuItemResponse{}, err
	}

	promotionMenuItemsResponse := make([]PromotionMenuItemResponse, len(prom.PromotionMenuItems))
	for i, promotionMenuItem := range prom.PromotionMenuItems {
		m, err := ps.menuRepo.GetById(promotionMenuItem.MenuItemId)

		if err != nil {
			return []PromotionMenuItemResponse{}, err
		}

		promotionMenuItemsResponse[i].Type = promotionMenuItem.Type
		promotionMenuItemsResponse[i].MenuItem = *m
	}

	return promotionMenuItemsResponse, nil
}
