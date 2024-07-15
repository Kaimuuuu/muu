package menu

import (
	"fmt"
	"github.com/Kaimuuuu/muu/model"
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
	promotions, err := ms.promotionRepo.All()
	if err != nil {
		return err
	}

	for _, p := range promotions {
		filtered := make([]model.PromotionMenuItem, 0)
		for _, pm := range p.PromotionMenuItems {
			if pm.MenuItemId != menuItemId {
				filtered = append(filtered, pm)
			}
		}
		p.PromotionMenuItems = filtered

		if err := ms.promotionRepo.Update(p.Id, &p); err != nil {
			return err
		}
	}

	return nil
}
