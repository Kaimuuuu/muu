package token

import (
	"kaimuu/model"
	"time"
)

func (ts *TokenService) Generate(req GenerateTokenRequest, employeeId string) (string, error) {
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
