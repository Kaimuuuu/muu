package order

import (
	"kaimuu/model"
	"time"

	"github.com/google/uuid"
)

func (os *OrderService) Create(req CreateOrderRequest, c *model.Client) error {
	p, err := os.promotionRepo.GetById(c.PromotionId)
	if err != nil {
		return err
	}

	orderItems := make([]model.OrderItem, len(req.OrderItems))
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

		if m.OutOfStock {
			return MenuItemOutOfStockError
		}

		price := m.Price
		for _, pmi := range p.PromotionMenuItems {
			if pmi.MenuItemId == roi.MenuItemId && pmi.Type == model.PromotionBuffet {
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

	o := &model.Order{
		Id:          uuid.NewString(),
		TableNumber: c.TableNumber,
		OrderItems:  orderItems,
		CreatedAt:   time.Now(),
		OrderBy:     c.Token,
	}

	if err := os.orderRepo.Insert(o); err != nil {
		return err
	}

	return nil
}
