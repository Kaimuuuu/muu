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

	for _, c := range clients {
		if c.TableNumber == req.TableNumber {
			return "", TableAlreadyInUsedError
		}
	}

	token := GenerateToken()

	p, err := cs.promotionServ.GetPromotionById(req.PromotionId)
	if err != nil {
		return "", err
	}

	c := &model.Client{
		TableNumber:   req.TableNumber,
		Size:          req.Size,
		PromotionId:   req.PromotionId,
		PromotionName: p.Name,
		Expire:        time.Now().Add(p.Duration),
		Token:         token,
		CreatedAt:     time.Now(),
		CreatedBy:     employeeId,
	}

	if err := cs.tokenStorage.Set(token, c); err != nil {
		return "", err
	}

	return token, err
}
