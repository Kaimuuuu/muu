package transaction

import (
	"kaimuu/model"
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
}

func (ts *TransactionService) toTransaction(c *model.Client) (*model.Transaction, error) {
	orders, err := ts.orderRepo.GetByToken(c.Token)
	if err != nil {
		return &model.Transaction{}, err
	}

	orderItems := make([]model.OrderItem, 0)
	for _, o := range orders {
		for _, oi := range o.OrderItems {
			if !oi.OutOfStock && oi.IsComplete {
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

	p, err := ts.promotionRepo.GetById(c.PromotionId)
	if err != nil {
		return &model.Transaction{}, err
	}

	sum += p.Price * float32(c.Size)
	for _, oi := range orderItems {
		sum += oi.Price * float32(oi.Quantity)
	}

	return &model.Transaction{
		TableNumber:       c.TableNumber,
		Size:              c.Size,
		PromotionName:     p.Name,
		TotalPrice:        sum,
		RemainingDuration: time.Until(c.Expire),
		CreatedAt:         c.CreatedAt,
		OrderItems:        orderItems,
	}, nil
}
