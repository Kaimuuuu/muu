package menu

import (
	"fmt"
	"kaimuu/model"
	"os"
)

func (ms *MenuService) Delete(menuItemId string) error {
	m, err := ms.menuRepo.GetById(menuItemId)
	if err != nil {
		return err
	}

	os.Remove(fmt.Sprintf("public/%s", m.ImagePath))

	if err := ms.menuRepo.Delete(menuItemId); err != nil {
		return err
	}

	// maybe move this logic to repository layer?
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
