package client

import (
	"kaimuu/model"
	"time"

	"github.com/google/uuid"
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
	Insert(trans *TransactionObject) error
}

type OrderService interface {
	GetOrderById(orderId string) (*model.Order, error)
	GetOrderByToken(token string) ([]model.Order, error)
}

type PromotionService interface {
	GetPromotionById(promotionId string) (*model.Promotion, error)
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

func (cs *ClientService) toTransactionObject(cli *model.Client) (*TransactionObject, error) {
	ol, err := cs.orderServ.GetOrderByToken(cli.Token)
	if err != nil {
		return &TransactionObject{}, err
	}

	orderItems := make([]model.OrderItem, 0)
	for _, o := range ol {
		if o.Status == model.Decline {
			continue
		}
		for _, oi := range o.OrderItems {
			if !oi.OutOfStock {
				orderItems = append(orderItems, oi)
			}
		}
	}

	collections := make(map[string]model.OrderItem)
	for _, oi := range orderItems {
		val, ok := collections[oi.MenuItemId]

		if !ok {
			collections[oi.MenuItemId] = oi
		} else {
			val.Quantity += oi.Quantity
			collections[oi.MenuItemId] = val
		}
	}

	summaryOrderItems := make([]model.OrderItem, 0)
	for _, oi := range collections {
		summaryOrderItems = append(summaryOrderItems, oi)
	}

	var sum float32 = 0

	promo, err := cs.promotionServ.GetPromotionById(cli.PromotionId)
	if err != nil {
		return &TransactionObject{}, err
	}

	sum += promo.Price
	for _, oi := range summaryOrderItems {
		sum += oi.Price * float32(oi.Quantity)
	}

	return &TransactionObject{
		TableNumber:       cli.TableNumber,
		Size:              cli.Size,
		PromotionName:     promo.Name,
		TotalPrice:        sum,
		RemainingDuration: time.Until(cli.Expire),
		CreatedAt:         cli.CreatedAt,
		OrderItems:        summaryOrderItems,
	}, nil
}
