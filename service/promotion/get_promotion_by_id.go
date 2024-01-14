package promotion

import "kaimuu/model"

func (ps *PromotionService) GetPromotionById(promotionId string) (*model.Promotion, error) {
	promo, err := ps.promotionRepo.GetById(promotionId)

	if err != nil {
		return &model.Promotion{}, err
	}

	return promo, nil
}
