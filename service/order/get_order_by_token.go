package order

import "kaimuu/model"

func (os *OrderService) GetOrderByToken(token string) ([]model.Order, error) {
	orders, err := os.orderRepo.GetOrderByToken(token)

	if err != nil {
		return []model.Order{}, err
	}

	return orders, nil
}
