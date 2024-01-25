package menu

import (
	"kaimuu/model"
)

func NewMenuService(menuRepo MenuRepository, promotionServ PromotionService, orderServ OrderService, promotionRepo PromotionRepository) *MenuService {
	return &MenuService{
		menuRepo:      menuRepo,
		promotionServ: promotionServ,
		orderServ:     orderServ,
		promotionRepo: promotionRepo,
	}
}

type MenuService struct {
	menuRepo      MenuRepository
	promotionServ PromotionService
	orderServ     OrderService
	promotionRepo PromotionRepository
}

type PromotionService interface {
	GetPromotionById(id string) (*model.Promotion, error)
	GetPromotions() ([]model.Promotion, error)
}

type OrderService interface {
	UpdateOutOfStockPendingOrder(menuItemId string, IsOutOfStock bool) error
}

type MenuRepository interface {
	Insert(m *model.MenuItem) error
	GetById(id string) (*model.MenuItem, error)
	GetAll() ([]model.MenuItem, error)
	Update(id string, m *model.MenuItem) error
	Delete(id string) error
}

type PromotionRepository interface {
	Update(id string, p *model.Promotion) error
}

type CreateMenuRequest struct {
	Name        string  `validate:"required" json:"name"`
	Catagory    string  `validate:"required" json:"catagory"`
	Weight      int     `validate:"required,number" json:"weight"`
	Description string  `validate:"" json:"description"`
	OutOfStock  bool    `validate:"boolean" json:"outOfStock"`
	Price       float32 `validate:"min=0,number" json:"price"`
	ImagePath   string  `validate:"" json:"imagePath"`
}

type UpdateMenuRequest struct {
	Name        string  `validate:"required" json:"name"`
	Catagory    string  `validate:"required" json:"catagory"`
	Weight      int     `validate:"required,number" json:"weight"`
	Description string  `validate:"" json:"description"`
	Price       float32 `validate:"min=0,number" json:"price"`
	ImagePath   string  `validate:"" json:"imagePath"`
}

type OutOfStockRequest struct {
	IsOutOfStock bool `validate:"boolean" json:"isOutOfStock"`
}

func New(menuRepo MenuRepository) *MenuService {
	return &MenuService{menuRepo: menuRepo}
}
