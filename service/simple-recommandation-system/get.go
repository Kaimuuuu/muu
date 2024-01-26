package simplerecommandationsystem

import (
	"kaimuu/model"
	"sort"
)

func (srs *SimpleRecommandationSystem) Get(promotionId string) ([]model.MenuItem, error) {
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

	p, err := srs.PromotionRepo.GetById(promotionId)
	if err != nil {
		return []model.MenuItem{}, err
	}

	filteredKey := make([]string, 0)
	for _, key := range keys {
		for _, promotionMenuItem := range p.PromotionMenuItems {
			if key == promotionMenuItem.MenuItemId {
				filteredKey = append(filteredKey, key)
			}
		}
	}

	menuItems := make([]model.MenuItem, 0)

	for _, menuItemId := range filteredKey {
		m, err := srs.MenuRepo.GetById(menuItemId)
		if err != nil {
			return []model.MenuItem{}, err
		}

		menuItems = append(menuItems, *m)
	}

	showLimit := 10

	if len(menuItems) < showLimit {
		return menuItems, nil
	}

	return menuItems[:showLimit], nil

}
