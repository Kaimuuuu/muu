package simplerecommandationsystem

import "kaimuu/model"

type SimpleRecommandationSystem struct {
	Recommands    map[string]int // key: MenuItemId, value: OrderCount
	MenuRepo      MenuRepository
	PromotionRepo PromotionRepository
}

type MenuRepository interface {
	All() ([]model.MenuItem, error)
	GetById(id string) (*model.MenuItem, error)
}

type PromotionRepository interface {
	GetById(id string) (*model.Promotion, error)
}

func New(menuRepo MenuRepository, promotionRepo PromotionRepository) *SimpleRecommandationSystem {
	return &SimpleRecommandationSystem{
		Recommands:    make(map[string]int),
		MenuRepo:      menuRepo,
		PromotionRepo: promotionRepo,
	}
}
