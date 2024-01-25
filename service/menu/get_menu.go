package menu

import "kaimuu/model"

func (ms *MenuService) GetMenu(cli *model.Client) ([]model.MenuItem, error) {
	p, err := ms.promotionServ.GetPromotionById(cli.PromotionId)
	if err != nil {
		return []model.MenuItem{}, err
	}

	menuItems := make([]model.MenuItem, 0)
	for _, pm := range p.PromotionMenuItems {
		m, err := ms.GetMenuItemById(pm.MenuItemId)
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
