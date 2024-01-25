package client

import (
	"kaimuu/model"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

var (
	TableAlreadyInUsedError = errors.New("table is already in used")
)

func NewClientService(transactionRepo TransactionRepository, orderServ OrderService, tokenStorage TokenStorage, promotionServ PromotionService) *ClientService {
	return &ClientService{
		transactionRepo: transactionRepo,
		orderServ:       orderServ,
		tokenStorage:    tokenStorage,
		promotionServ:   promotionServ,
	}
}

type ClientService struct {
	transactionRepo TransactionRepository
	orderServ       OrderService
	tokenStorage    TokenStorage
	promotionServ   PromotionService
}

type TransactionRepository interface {
	Insert(t *TransactionObject) error
}

type OrderService interface {
	GetOrderById(id string) (*model.Order, error)
	GetOrderByToken(token string) ([]model.Order, error)
}

type PromotionService interface {
	GetPromotionById(id string) (*model.Promotion, error)
}

type TokenStorage interface {
	Get(key string) (*model.Client, error)
	Set(key string, value *model.Client) error
	Remove(key string) error
	GetAll() ([]model.Client, error)
}

func GenerateToken() string {
	return uuid.NewString()
}

type SignInRequest struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type GenerateClientRequest struct {
	PromotionId string `validate:"required" json:"promotionId"`
	TableNumber int8   `validate:"required,number" json:"tableNumber"`
	Size        int8   `validate:"min=1" json:"size"`
}

type TransactionObject struct {
	TableNumber       int8              `json:"tableNumber"`
	Size              int8              `json:"size"`
	PromotionName     string            `json:"promotionName"`
	TotalPrice        float32           `json:"totalPrice"`
	RemainingDuration time.Duration     `json:"remainingDuration"`
	CreatedAt         time.Time         `json:"createdAt"`
	OrderItems        []model.OrderItem `json:"orderItems"`
}

func (cs *ClientService) toTransactionObject(c *model.Client) (*TransactionObject, error) {
	orders, err := cs.orderServ.GetOrderByToken(c.Token)
	if err != nil {
		return &TransactionObject{}, err
	}

	orderItems := make([]model.OrderItem, 0)
	for _, o := range orders {
		if o.Status == model.Decline {
			continue
		}
		for _, oi := range o.OrderItems {
			if !oi.OutOfStock {
				orderItems = append(orderItems, oi)
			}
		}
	}

	col := make(map[string]model.OrderItem)
	for _, oi := range orderItems {
		val, ok := col[oi.MenuItemId]

		if !ok {
			col[oi.MenuItemId] = oi
		} else {
			val.Quantity += oi.Quantity
			col[oi.MenuItemId] = val
		}
	}

	orderItems = make([]model.OrderItem, 0)
	for _, oi := range col {
		orderItems = append(orderItems, oi)
	}

	var sum float32 = 0

	p, err := cs.promotionServ.GetPromotionById(c.PromotionId)
	if err != nil {
		return &TransactionObject{}, err
	}

	sum += p.Price * float32(c.Size)
	for _, oi := range orderItems {
		sum += oi.Price * float32(oi.Quantity)
	}

	return &TransactionObject{
		TableNumber:       c.TableNumber,
		Size:              c.Size,
		PromotionName:     p.Name,
		TotalPrice:        sum,
		RemainingDuration: time.Until(c.Expire),
		CreatedAt:         c.CreatedAt,
		OrderItems:        orderItems,
	}, nil
}
