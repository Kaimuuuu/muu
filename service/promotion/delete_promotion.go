package promotion

import (
	"fmt"
	"os"
)

func (ps *PromotionService) Delete(promotionId string) error {
	clients, err := ps.tokenStorage.GetAll()
	if err != nil {
		return err
	}

	for _, cli := range clients {
		if cli.PromotionId == promotionId {
			return ClientInUsedError
		}
	}

	promo, err := ps.promotionRepo.GetById(promotionId)
	if err != nil {
		return err
	}

	os.Remove(fmt.Sprintf("public/%s", promo.ImagePath))

	if err := ps.promotionRepo.Delete(promotionId); err != nil {
		return err
	}

	return nil
}
