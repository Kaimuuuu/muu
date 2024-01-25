package order

import (
	"kaimuu/model"
	"time"

	"github.com/google/uuid"
)

func (os *OrderService) CreateOrder(req CreateOrderRequest, c *model.Client) error {
	p, err := os.promotionServ.GetPromotionById(c.PromotionId)
	if err != nil {
		return err
	}

	orderItems := make([]model.OrderItem, len(req.OrderItems))

	totalWeight := 0
	for i, roi := range req.OrderItems {
		m, err := os.menuRepo.GetById(roi.MenuItemId)
		if err != nil {
			return err
		}

		// validate menuItemId
		errBit := 1
		for _, promotionMenuItem := range p.PromotionMenuItems {
			if promotionMenuItem.MenuItemId == roi.MenuItemId {
				errBit = 0
				break
			}
		}
		if errBit == 1 {
			return OrderInvalidMenuItemError
		}

		totalWeight += m.Weight * int(roi.Quantity)

		if m.OutOfStock {
			return MenuItemOutOfStockError
		}

		price := m.Price
		for _, promotionMenuItem := range p.PromotionMenuItems {
			if promotionMenuItem.MenuItemId == roi.MenuItemId && promotionMenuItem.Type == model.PromotionBuffet {
				price = 0
			}
		}

		orderItems[i] = model.OrderItem{
			MenuItemId: roi.MenuItemId,
			Quantity:   roi.Quantity,
			Name:       m.Name,
			OutOfStock: m.OutOfStock,
			Price:      price,
		}
	}

	if totalWeight > p.Weight {
		return WeightExceededError
	}

	o := &model.Order{
		Id:          uuid.NewString(),
		TableNumber: c.TableNumber,
		OrderItems:  orderItems,
		Status:      model.Pending,
		CreatedAt:   time.Now(),
		OrderBy:     c.Token,
	}

	if err := os.orderRepo.Insert(o); err != nil {
		return err
	}

	return nil
}
