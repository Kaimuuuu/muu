package promotion

import (
	"kaimuu/model"
	"time"

	"github.com/google/uuid"
)

func (ps *PromotionService) Create(req CreatePromotionRequest, employeeId string) error {
	p := &model.Promotion{
		Id:                 uuid.NewString(),
		Name:               req.Name,
		Description:        req.Description,
		Price:              req.Price,
		Duration:           req.Duration,
		PromotionMenuItems: req.PromotionMenuItems,
		ImagePath:          req.ImagePath,
		CreatedAt:          time.Now(),
		CreatedBy:          employeeId,
	}

	if err := ps.promotionRepo.Insert(p); err != nil {
		return err
	}

	return nil
}
