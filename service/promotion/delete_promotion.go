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

	if err := ps.promotionRepo.Delete(promotionId); err != nil {
		return err
	}

	return nil
}
