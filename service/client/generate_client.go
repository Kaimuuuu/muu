package client

import (
	"kaimuu/model"
	"time"
)

func (cs *ClientService) GenerateClient(req GenerateClientRequest, employeeId string) (string, error) {
	clients, err := cs.tokenStorage.GetAll()
	if err != nil {
		return "", err
	}

	for _, client := range clients {
		if client.TableNumber == req.TableNumber {
			return "", TableAlreadyInUsedError
		}
	}

	token := GenerateToken()

	promo, err := cs.promotionServ.GetPromotionById(req.PromotionId)
	if err != nil {
		return "", err
	}

	cli := &model.Client{
		TableNumber:   req.TableNumber,
		Size:          req.Size,
		PromotionId:   req.PromotionId,
		PromotionName: promo.Name,
		Expire:        time.Now().Add(promo.Duration),
		Token:         token,
		CreatedAt:     time.Now(),
		CreatedBy:     employeeId,
	}

	if err := cs.tokenStorage.Set(token, cli); err != nil {
		return "", err
	}

	return token, err
}
