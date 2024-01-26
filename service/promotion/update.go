package promotion

import (
	"fmt"
	"os"
)

func (ps *PromotionService) Update(promotionId string, req UpdatePromotionRequest) error {
	clients, err := ps.tokenRepo.All()
	if err != nil {
		return err
	}

	for _, c := range clients {
		if c.PromotionId == promotionId {
			return ClientInUsedError
		}
	}

	p, err := ps.promotionRepo.GetById(promotionId)
	if err != nil {
		return err
	}

	os.Remove(fmt.Sprintf("public/%s", p.ImagePath))

	p.Name = req.Name
	p.Description = req.Description
	p.Price = req.Price
	p.Duration = req.Duration
	p.PromotionMenuItems = req.PromotionMenuItems
	p.ImagePath = req.ImagePath

	if err := ps.promotionRepo.Update(promotionId, p); err != nil {
		return err
	}

	return nil
}
