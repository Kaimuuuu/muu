package menu

import (
	"github.com/Kaimuuuu/muu/model"
	"github.com/Kaimuuuu/muu/service/promotion"
)

func NewMenuService(menuRepo MenuRepository, orderServ OrderService, promotionRepo PromotionRepository, promotionServ PromotionService) *MenuService {
	return &MenuService{
		menuRepo:      menuRepo,
		orderServ:     orderServ,
		promotionRepo: promotionRepo,
		promotionServ: promotionServ,
	}
}

type MenuService struct {
	menuRepo      MenuRepository
	promotionRepo PromotionRepository
	orderServ     OrderService
	promotionServ PromotionService
}

type PromotionRepository interface {
	GetById(id string) (*model.Promotion, error)
	All() ([]model.Promotion, error)
	Update(id string, p *model.Promotion) error
}

type PromotionService interface {
	GetMenu(promotionId string) ([]promotion.PromotionMenuItemResponse, error)
}

type OrderService interface {
	UpdateOutOfStockPendingOrders(menuItemId string, IsOutOfStock bool) error
}

type MenuRepository interface {
	Insert(m *model.MenuItem) error
	GetById(id string) (*model.MenuItem, error)
	All() ([]model.MenuItem, error)
	Update(id string, m *model.MenuItem) error
	Delete(id string) error
}

type CreateMenuRequest struct {
	Name        string  `validate:"required" json:"name"`
	Catagory    string  `validate:"required" json:"catagory"`
	Description string  `validate:"" json:"description"`
	OutOfStock  bool    `validate:"boolean" json:"outOfStock"`
	Price       float32 `validate:"min=0,number" json:"price"`
	ImagePath   string  `validate:"" json:"imagePath"`
}

type UpdateMenuRequest struct {
	Name        string  `validate:"required" json:"name"`
	Catagory    string  `validate:"required" json:"catagory"`
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
