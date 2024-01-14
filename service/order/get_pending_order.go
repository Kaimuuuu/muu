package order

import "kaimuu/model"

func (os *OrderService) GetPendingOrder() ([]model.Order, error) {
	ol, err := os.orderRepo.GetPendingOrder()

	if err != nil {
		return []model.Order{}, err
	}

	return ol, nil
}
