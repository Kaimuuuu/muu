package simplerecommandationsystem

import (
	"kaimuu/model"
	"sort"
)

func (srs *SimpleRecommandationSystem) GetRecommand(promotionId string) ([]model.MenuItem, error) {
	if err := srs.Sync(); err != nil {
		return []model.MenuItem{}, err
	}

	keys := make([]string, 0)

	for key := range srs.Recommands {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return srs.Recommands[keys[i]] > srs.Recommands[keys[j]]
	})

	promo, err := srs.PromotionRepo.GetById(promotionId)
	if err != nil {
		return []model.MenuItem{}, err
	}

	filteredKey := make([]string, 0)
	for _, key := range keys {
		for _, promotionMenuItem := range promo.PromotionMenuItems {
			if key == promotionMenuItem.MenuItemId {
				filteredKey = append(filteredKey, key)
			}
		}
	}

	menus := make([]model.MenuItem, 0)

	for _, menuItemId := range filteredKey {
		menu, err := srs.MenuRepo.GetById(menuItemId)
		if err != nil {
			return []model.MenuItem{}, err
		}

		menus = append(menus, *menu)
	}

	return menus, nil
}
