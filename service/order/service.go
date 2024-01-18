package order

import (
	"kaimuu/model"

	"github.com/cockroachdb/errors"
)

var (
	OrderInvalidMenuItemError = errors.New("ordering invalid menu item id")
	MenuItemOutOfStockError   = errors.New("menu item out of stock")
	WeightExceededError       = errors.New("total weight exceeded")
)

func NewOrderService(orderRepo OrderRepository, menuRepo MenuRepository, promotionServ PromotionService, tokenStorage TokenStorage) *OrderService {
	return &OrderService{
		orderRepo:     orderRepo,
		menuRepo:      menuRepo,
		promotionServ: promotionServ,
		tokenStorage:  tokenStorage,
	}
}

type OrderService struct {
	orderRepo     OrderRepository
	menuRepo      MenuRepository
	promotionServ PromotionService
	tokenStorage  TokenStorage
}

type MenuRepository interface {
	GetById(id string) (*model.MenuItem, error)
}

type PromotionService interface {
	GetPromotionById(promotionId string) (*model.Promotion, error)
}

type OrderRepository interface {
	Insert(o *model.Order) error
	GetById(id string) (*model.Order, error)
	Delete(id string) error
	GetPendingOrder() ([]model.Order, error)
	GetOrderByToken(token string) ([]model.Order, error)
	Update(id string, o *model.Order) error
}

type TokenStorage interface {
	Set(key string, cli *model.Client) error
}

type CreateOrderRequest struct {
	OrderItems []RequestOrderItem `validate:"gt=0,required,dive" json:"orderItems"`
}

type RequestOrderItem struct {
	MenuItemId string `validate:"required" json:"menuItemId"`
	Quantity   int8   `validate:"required,number" json:"quantity"`
}

type UpdateOrderStatusRequest struct {
	Status model.OrderStatus `validate:"required" json:"status"`
}
