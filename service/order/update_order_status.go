package order

import "kaimuu/model"

func (os *OrderService) UpdateOrderStatus(req UpdateOrderStatusRequest, orderId string) error {
	o, err := os.orderRepo.GetById(orderId)
	if err != nil {
		return err
	}

	o.Status = req.Status

	if req.Status == model.OrderSuccessStatus {
		for i, oi := range o.OrderItems {
			if !oi.OutOfStock {
				o.OrderItems[i].IsComplete = true
			}
		}
	}

	if err := os.orderRepo.Update(orderId, o); err != nil {
		return err
	}

	return nil
}
