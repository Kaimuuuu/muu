package order

func (os *OrderService) UpdateOrderItemStatus(req UpdateOrderItemsStatusRequest, orderId string) error {
	o, err := os.orderRepo.GetById(orderId)
	if err != nil {
		return err
	}

	for i, oi := range o.OrderItems {
		for _, roi := range req.OrderItemsStatus {
			if oi.MenuItemId == roi.MenuItemId {
				o.OrderItems[i].IsComplete = roi.Status
			}
		}
	}

	if err := os.orderRepo.Update(orderId, o); err != nil {
		return err
	}

	return nil
}
