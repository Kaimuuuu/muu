package order

import "kaimuu/model"

func (os *OrderService) GetOrderById(orderId string) (*model.Order, error) {
	o, err := os.orderRepo.GetById(orderId)

	if err != nil {
		return &model.Order{}, err
	}

	return o, nil
}
