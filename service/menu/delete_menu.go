package menu

func (ms *MenuService) Delete(menuItemId string) error {
	if err := ms.menuRepo.Delete(menuItemId); err != nil {
		return err
	}

	return nil
}
