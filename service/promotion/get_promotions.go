package promotion

import "kaimuu/model"

func (ps *PromotionService) GetPromotions() ([]model.Promotion, error) {
	promotions, err := ps.promotionRepo.GetAll()

	if err != nil {
		return []model.Promotion{}, err
	}

	return promotions, nil
}
