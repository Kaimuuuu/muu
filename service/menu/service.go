package menu

import (
	"kaimuu/model"
)

func NewMenuService(menuRepo MenuRepository, promotionServ PromotionService, orderServ OrderService) *MenuService {
	return &MenuService{
		menuRepo:      menuRepo,
		promotionServ: promotionServ,
		orderServ:     orderServ,
	}
}

type MenuService struct {
	menuRepo      MenuRepository
	promotionServ PromotionService
	orderServ     OrderService
}

type PromotionService interface {
	GetPromotionById(promotionId string) (*model.Promotion, error)
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

type CreateMenuRequest struct {
	Name        string  `validate:"required" json:"name"`
	Catagory    string  `validate:"required" json:"catagory"`
	Weight      int     `validate:"required,number" json:"weight"`
	Description string  `validate:"required" json:"description"`
	OutOfStock  bool    `validate:"boolean" json:"outOfStock"`
	Price       float32 `validate:"required,number" json:"price"`
	ImagePath   string  `validate:"imagePath" json:"imagePath"`
}

type UpdateMenuRequest struct {
	Name        string  `validate:"required" json:"name"`
	Catagory    string  `validate:"required" json:"catagory"`
	Weight      int     `validate:"required,number" json:"weight"`
	Description string  `validate:"required" json:"description"`
	Price       float32 `validate:"required,number" json:"price"`
	ImagePath   string  `validate:"imagePath" json:"imagePath"`
}

type OutOfStockRequest struct {
	IsOutOfStock bool `validate:"boolean" json:"isOutOfStock"`
}

func New(menuRepo MenuRepository) *MenuService {
	return &MenuService{menuRepo: menuRepo}
}
