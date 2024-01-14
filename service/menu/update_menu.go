package menu

func (ms *MenuService) UpdateMenu(menuItemId string, req UpdateMenuRequest) error {
	m, err := ms.GetMenuItemById(menuItemId)

	if err != nil {
		return err
	}

	m.Catagory = req.Catagory
	m.Description = req.Description
	m.Name = req.Name
	m.Price = req.Price
	m.Weight = req.Weight
	m.ImagePath = req.ImagePath

	if err := ms.menuRepo.Update(menuItemId, m); err != nil {
		return err
	}

	return nil
}
