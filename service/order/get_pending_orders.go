package order

import "kaimuu/model"

func (os *OrderService) GetPendingOrders() ([]model.Order, error) {
	orders, err := os.orderRepo.GetPendingOrders()
	if err != nil {
		return []model.Order{}, err
	}

	return orders, nil
}
