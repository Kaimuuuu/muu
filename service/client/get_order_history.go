package client

import "kaimuu/model"

func (cs *ClientService) GetOrderHistory(cli *model.Client) ([]model.Order, error) {
	ol, err := cs.orderServ.GetOrderByToken(cli.Token)

	if err != nil {
		return []model.Order{}, err
	}

	return ol, nil
}
