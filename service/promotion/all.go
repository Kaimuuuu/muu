package promotion

import "github.com/Kaimuuuu/muu/model"

func (ps *PromotionService) All() ([]model.Promotion, error) {
	promotions, err := ps.promotionRepo.All()
	if err != nil {
		return []model.Promotion{}, err
	}

	return promotions, nil
}
