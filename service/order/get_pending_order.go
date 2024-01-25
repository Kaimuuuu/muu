package order

import "kaimuu/model"

func (os *OrderService) GetPendingOrder() ([]model.Order, error) {
	orders, err := os.orderRepo.GetPendingOrder()
	if err != nil {
		return []model.Order{}, err
	}

	return orders, nil
}
