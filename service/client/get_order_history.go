package client

import "kaimuu/model"

func (cs *ClientService) GetOrderHistory(c *model.Client) ([]model.Order, error) {
	orders, err := cs.orderServ.GetOrderByToken(c.Token)

	if err != nil {
		return []model.Order{}, err
	}

	return orders, nil
}
