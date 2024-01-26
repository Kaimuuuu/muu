package order

func (os *OrderService) Delete(orderId string) error {
	if err := os.orderRepo.Delete(orderId); err != nil {
		return err
	}

	return nil
}
