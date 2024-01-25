package menu

func (ms *MenuService) SetOutOfStock(menuItemId string, isOutOfStock bool) error {
	m, err := ms.menuRepo.GetById(menuItemId)
	if err != nil {
		return err
	}

	m.OutOfStock = isOutOfStock

	if err := ms.menuRepo.Update(menuItemId, m); err != nil {
		return err
	}

	if err := ms.orderServ.UpdateOutOfStockPendingOrder(menuItemId, isOutOfStock); err != nil {
		return err
	}

	return nil
}
