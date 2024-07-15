package transaction

import (
	"fmt"
	"github.com/Kaimuuuu/muu/model"
	"time"
)

func NewTransactionService(transactionRepo TransactionRepository, orderRepo OrderRepository, promotionRepo PromotionRepository, tokenRepo TokenRepository) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
		orderRepo:       orderRepo,
		promotionRepo:   promotionRepo,
		tokenRepo:       tokenRepo,
	}
}

type TransactionService struct {
	transactionRepo TransactionRepository
	orderRepo       OrderRepository
	promotionRepo   PromotionRepository
	tokenRepo       TokenRepository
}

type OrderRepository interface {
	GetById(id string) (*model.Order, error)
	GetByToken(token string) ([]model.Order, error)
}

type PromotionRepository interface {
	GetById(id string) (*model.Promotion, error)
}

type TokenRepository interface {
	Get(token string) (*model.Client, error)
	Delete(token string) error
}

type TransactionRepository interface {
	Insert(t *model.Transaction) error
	All() ([]model.Transaction, error)
}

func (ts *TransactionService) toTransaction(c *model.Client) (*model.Transaction, error) {
	orders, err := ts.orderRepo.GetByToken(c.Token)
	if err != nil {
		return &model.Transaction{}, err
	}

	orderItems := make([]model.OrderItem, 0)
	for _, o := range orders {
		for _, oi := range o.OrderItems {
			if oi.IsComplete || (o.Status == model.OrderPendingStatus && !oi.IsComplete) {
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

	var startingPrice float32 = 0
	var orderPrice float32 = 0

	p, err := ts.promotionRepo.GetById(c.PromotionId)
	if err != nil {
		return &model.Transaction{}, err
	}

	startingPrice += p.Price * float32(c.Size)
	for _, oi := range orderItems {
		orderPrice += oi.Price * float32(oi.Quantity)
	}

	return &model.Transaction{
		TableNumber:       c.TableNumber,
		Token:             c.Token,
		Size:              c.Size,
		PromotionName:     fmt.Sprintf("%s (%.2f บาท)", p.Name, p.Price),
		StartPrice:        startingPrice,
		OrderPrice:        orderPrice,
		TotalPrice:        startingPrice + orderPrice,
		RemainingDuration: time.Until(c.Expire),
		CreatedAt:         c.CreatedAt,
		OrderItems:        orderItems,
	}, nil
}
