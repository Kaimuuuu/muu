package order

import (
	"kaimuu/model"

	"github.com/cockroachdb/errors"
)

var (
	OrderInvalidMenuItemError = errors.New("ordering invalid menu item id")
	MenuItemOutOfStockError   = errors.New("menu item out of stock")
	WeightExceededError       = errors.New("total weight exceeded")
	InvalidOrderQuantity      = errors.New("order request exceed promotion limit")
)

func NewOrderService(orderRepo OrderRepository, menuRepo MenuRepository, promotionRepo PromotionRepository, tokenRepo TokenRepository) *OrderService {
	return &OrderService{
		orderRepo:     orderRepo,
		menuRepo:      menuRepo,
		promotionRepo: promotionRepo,
		tokenRepo:     tokenRepo,
	}
}

type OrderService struct {
	orderRepo     OrderRepository
	menuRepo      MenuRepository
	promotionRepo PromotionRepository
	tokenRepo     TokenRepository
}

type MenuRepository interface {
	GetById(id string) (*model.MenuItem, error)
}

type PromotionRepository interface {
	GetById(promotionId string) (*model.Promotion, error)
}

type OrderRepository interface {
	Insert(o *model.Order) error
	GetById(id string) (*model.Order, error)
	Delete(id string) error
	GetPendingOrders() ([]model.Order, error)
	GetByToken(token string) ([]model.Order, error)
	Update(id string, o *model.Order) error
}

type TokenRepository interface {
	Insert(c *model.Client) error
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

type UpdateOrderItemsStatusRequest struct {
	OrderItemsStatus []RequestOrderItemStatus `validate:"gt=0,required,dive" json:"orderItemsStatus"`
}

type RequestOrderItemStatus struct {
	MenuItemId string `validate:"required" json:"menuItemId"`
	Status     bool   `validate:"boolean" json:"status"`
}
