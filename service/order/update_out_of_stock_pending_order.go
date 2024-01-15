package order

func (os *OrderService) UpdateOutOfStockPendingOrder(menuItemId string, isOutOfStock bool) error {
	// Time Complexity: O(o_p * c) ; o_p is size of pending order, c is size of menu
	op, err := os.GetPendingOrder()

	if err != nil {
		return err
	}

	for _, o := range op {
		for i, oi := range o.OrderItems {
			if oi.MenuItemId == menuItemId {
				o.OrderItems[i].OutOfStock = isOutOfStock
			}
		}
		if err := os.orderRepo.Update(o.Id, &o); err != nil {
			return err
		}
	}

	return nil
}
