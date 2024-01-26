package promotion

import "kaimuu/model"

func (ps *PromotionService) All() ([]model.Promotion, error) {
	promotions, err := ps.promotionRepo.All()
	if err != nil {
		return []model.Promotion{}, err
	}

	return promotions, nil
}
