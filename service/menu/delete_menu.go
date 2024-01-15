package menu

import "kaimuu/model"

func (ms *MenuService) Delete(menuItemId string) error {
	if err := ms.menuRepo.Delete(menuItemId); err != nil {
		return err
	}

	promotions, err := ms.promotionServ.GetPromotions()
	if err != nil {
		return err
	}

	for _, promo := range promotions {
		filtered := make([]model.PromotionMenuItem, 0)
		for _, promotionMenuItem := range promo.PromotionMenuItems {
			if promotionMenuItem.MenuItemId != menuItemId {
				filtered = append(filtered, promotionMenuItem)
			}
		}
		promo.PromotionMenuItems = filtered

		if err := ms.promotionRepo.Update(promo.Id, &promo); err != nil {
			return err
		}
	}

	return nil
}
