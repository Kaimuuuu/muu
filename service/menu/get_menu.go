package menu

import "kaimuu/model"

func (ms *MenuService) GetMenu(cli *model.Client) ([]model.MenuItem, error) {
	prom, err := ms.promotionServ.GetPromotionById(cli.PromotionId)

	if err != nil {
		return []model.MenuItem{}, err
	}

	menu := make([]model.MenuItem, 0)
	for _, promotionMenuItem := range prom.PromotionMenuItems {
		m, err := ms.GetMenuItemById(promotionMenuItem.MenuItemId)
		if err != nil {
			return []model.MenuItem{}, err
		}

		if promotionMenuItem.Type == model.Buffet {
			m.Price = 0
		}

		menu = append(menu, *m)
	}

	return menu, nil
}
