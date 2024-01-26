package promotion

import (
	"kaimuu/model"
	"time"

	"github.com/cockroachdb/errors"
)

var (
	ClientInUsedError = errors.New("there is some client used this promotion")
)

func NewPromotionService(promotionRepo PromotionRepository, menuRepo MenuRepository, tokenRepo TokenRepository) *PromotionService {
	return &PromotionService{
		promotionRepo: promotionRepo,
		menuRepo:      menuRepo,
		tokenRepo:     tokenRepo,
	}
}

type PromotionService struct {
	promotionRepo PromotionRepository
	menuRepo      MenuRepository
	tokenRepo     TokenRepository
}

type PromotionRepository interface {
	Insert(p *model.Promotion) error
	GetById(id string) (*model.Promotion, error)
	All() ([]model.Promotion, error)
	Update(id string, p *model.Promotion) error
	Delete(id string) error
}

type MenuRepository interface {
	GetById(id string) (*model.MenuItem, error)
}

type TokenRepository interface {
	Get(token string) (*model.Client, error)
	All() ([]model.Client, error)
	Delete(token string) error
	Insert(c *model.Client) error
}

type CreatePromotionRequest struct {
	Name               string                    `validate:"required" json:"name"`
	Price              float32                   `validate:"min=0,number" json:"price"`
	ImagePath          string                    `validate:"" json:"imagePath"`
	Duration           time.Duration             `validate:"required" json:"duration"`
	Description        string                    `validate:"" json:"description"`
	PromotionMenuItems []model.PromotionMenuItem `validate:"required,dive,required" json:"promotionMenuItems" form:"promotionMenuItems[]"`
}

type UpdatePromotionRequest struct {
	Name               string                    `validate:"required" json:"name"`
	Price              float32                   `validate:"min=0,number" json:"price"`
	ImagePath          string                    `validate:"" json:"imagePath"`
	Duration           time.Duration             `validate:"required" json:"duration"`
	Description        string                    `validate:"" json:"description"`
	PromotionMenuItems []model.PromotionMenuItem `validate:"required,dive,required" json:"promotionMenuItems" form:"promotionMenuItems[]"`
}

type PromotionMenuItemResponse struct {
	Type     model.PromotionMenuItemType `json:"type"`
	MenuItem model.MenuItem              `json:"menuItem"`
}
