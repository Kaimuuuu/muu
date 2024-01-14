package order

func (os *OrderService) UpdateOrderStatus(req UpdateOrderStatusRequest, orderId string) error {
	o, err := os.orderRepo.GetById(orderId)

	if err != nil {
		return err
	}

	o.Status = req.Status

	if err := os.orderRepo.Update(orderId, o); err != nil {
		return err
	}

	return nil
}
