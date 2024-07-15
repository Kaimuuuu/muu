package menu

import "github.com/Kaimuuuu/muu/model"

func (ms *MenuService) All() ([]model.MenuItem, error) {
	menus, err := ms.menuRepo.All()
	if err != nil {
		return []model.MenuItem{}, err
	}

	return menus, nil
}
