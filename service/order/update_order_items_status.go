package order

import "kaimuu/model"

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

	errBit := 0
	for _, oi := range o.OrderItems {
		if (!oi.IsComplete && !oi.OutOfStock) || (oi.IsComplete && oi.OutOfStock) {
			errBit = 1
		}
	}
	if errBit == 0 {
		o.Status = model.OrderSuccessStatus
	}

	if err := os.orderRepo.Update(orderId, o); err != nil {
		return err
	}

	return nil
}
