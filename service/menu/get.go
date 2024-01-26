package menu

import "kaimuu/model"

func (ms *MenuService) Get(cli *model.Client) ([]model.MenuItem, error) {
	p, err := ms.promotionRepo.GetById(cli.PromotionId)
	if err != nil {
		return []model.MenuItem{}, err
	}

	menuItems := make([]model.MenuItem, 0)
	for _, pm := range p.PromotionMenuItems {
		m, err := ms.menuRepo.GetById(pm.MenuItemId)
		if err != nil {
			return []model.MenuItem{}, err
		}

		if pm.Type == model.PromotionBuffet {
			m.Price = 0
		}

		menuItems = append(menuItems, *m)
	}

	return menuItems, nil
}
