package simplerecommandationsystem

import "kaimuu/model"

type SimpleRecommandationSystem struct {
	Recommands    map[string]int // key: MenuItemId, value: OrderCount
	MenuRepo      MenuRepository
	PromotionRepo PromotionRepository
}

type MenuRepository interface {
	GetAll() ([]model.MenuItem, error)
	GetById(menuItemId string) (*model.MenuItem, error)
}

type PromotionRepository interface {
	GetById(promotionId string) (*model.Promotion, error)
}

func New(menuRepo MenuRepository, promotionRepo PromotionRepository) *SimpleRecommandationSystem {
	return &SimpleRecommandationSystem{
		Recommands:    make(map[string]int),
		MenuRepo:      menuRepo,
		PromotionRepo: promotionRepo,
	}
}
