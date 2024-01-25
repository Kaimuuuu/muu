package promotion

import "kaimuu/model"

func (ps *PromotionService) GetPromotionById(promotionId string) (*model.Promotion, error) {
	p, err := ps.promotionRepo.GetById(promotionId)
	if err != nil {
		return &model.Promotion{}, err
	}

	return p, nil
}
