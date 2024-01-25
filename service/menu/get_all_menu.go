package menu

import "kaimuu/model"

func (ms *MenuService) GetAllMenu() ([]model.MenuItem, error) {
	menus, err := ms.menuRepo.GetAll()
	if err != nil {
		return []model.MenuItem{}, err
	}

	return menus, nil
}
