package promotion

import (
	"kaimuu/model"
	"time"
)

func NewPromotionService(promotionRepo PromotionRepository, menuRepo MenuRepository) *PromotionService {
	return &PromotionService{
		promotionRepo: promotionRepo,
		menuRepo:      menuRepo,
	}
}

type PromotionService struct {
	promotionRepo PromotionRepository
	menuRepo      MenuRepository
}

type PromotionRepository interface {
	Insert(promo *model.Promotion) error
	GetById(id string) (*model.Promotion, error)
	GetAll() ([]model.Promotion, error)
	Update(id string, promo *model.Promotion) error
	Delete(id string) error
}

type MenuRepository interface {
	GetById(menuItemId string) (*model.MenuItem, error)
}

type CreatePromotionRequest struct {
	Name               string                    `validate:"required" json:"name"`
	Weight             int                       `validate:"required,number" json:"weight"`
	Price              float32                   `validate:"required,number" json:"price"`
	ImagePath          string                    `validate:"" json:"imagePath"`
	Duration           time.Duration             `validate:"required" json:"duration"`
	Description        string                    `validate:"" json:"description"`
	PromotionMenuItems []model.PromotionMenuItem `validate:"required,dive,required" json:"promotionMenuItems" form:"promotionMenuItems[]"`
}

type UpdatePromotionRequest struct {
	Name               string                    `validate:"required" json:"name"`
	Weight             int                       `validate:"required,number" json:"weight"`
	Price              float32                   `validate:"required,number" json:"price"`
	ImagePath          string                    `validate:"" json:"imagePath"`
	Duration           time.Duration             `validate:"required" json:"duration"`
	Description        string                    `validate:"" json:"description"`
	PromotionMenuItems []model.PromotionMenuItem `validate:"required,dive,required" json:"promotionMenuItems" form:"promotionMenuItems[]"`
}

type UpdateWeightRequest struct {
	Weight int `validate:"required,number" json:"weight"`
}

type PromotionMenuItemResponse struct {
	Type     model.PromotionMenuItemType `json:"type"`
	MenuItem model.MenuItem              `json:"menuItem"`
}
