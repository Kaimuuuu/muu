package menu

import "kaimuu/model"

func (ms *MenuService) GetMenuItemById(menuItemId string) (*model.MenuItem, error) {
	m, err := ms.menuRepo.GetById(menuItemId)

	if err != nil {
		return &model.MenuItem{}, err
	}

	return m, nil
}
