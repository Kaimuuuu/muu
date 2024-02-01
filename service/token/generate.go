package token

import (
	"kaimuu/model"
	"time"
)

func (ts *TokenService) Generate(req GenerateTokenRequest, employeeId string) (string, error) {
	clients, err := ts.tokenRepo.All()
	if err != nil {
		return "", err
	}

	for _, c := range clients {
		if c.TableNumber == req.TableNumber {
			return "", TableInUsedError
		}
	}

	token := GenerateToken()

	p, err := ts.promotionRepo.GetById(req.PromotionId)
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

	if err := ts.tokenRepo.Insert(c); err != nil {
		return "", err
	}

	return token, err
}
