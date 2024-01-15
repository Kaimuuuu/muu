package order

import (
	"fmt"
	"kaimuu/model"
	"time"

	"github.com/google/uuid"
)

func (os *OrderService) CreateOrder(req CreateOrderRequest, cli *model.Client) error {
	promo, err := os.promotionServ.GetPromotionById(cli.PromotionId)

	if err != nil {
		return err
	}

	oi := make([]model.OrderItem, len(req.OrderItems))

	totalWeight := 0
	for i, roi := range req.OrderItems {
		menu, err := os.menuRepo.GetById(roi.MenuItemId)
		if err != nil {
			return err
		}

		// validate menuItemId
		errBit := 1
		for _, promotionMenuItem := range promo.PromotionMenuItems {
			if promotionMenuItem.MenuItemId == roi.MenuItemId {
				errBit = 0
				break
			}
		}
		if errBit == 1 {
			return fmt.Errorf("order invalid menu {%s}", menu.Name)
		}

		totalWeight += menu.Weight * int(roi.Quantity)

		if menu.OutOfStock {
			return fmt.Errorf("menu {%s} is out of stock", menu.Name)
		}

		price := menu.Price
		for _, promotionMenuItem := range promo.PromotionMenuItems {
			if promotionMenuItem.MenuItemId == roi.MenuItemId && promotionMenuItem.Type == model.Buffet {
				price = 0
			}
		}

		oi[i] = model.OrderItem{
			MenuItemId: roi.MenuItemId,
			Quantity:   roi.Quantity,
			Name:       menu.Name,
			OutOfStock: menu.OutOfStock,
			Price:      price,
		}
	}

	if totalWeight > promo.Weight {
		return fmt.Errorf("total weight exceeded")
	}

	o := &model.Order{
		Id:          uuid.NewString(),
		TableNumber: cli.TableNumber,
		OrderItems:  oi,
		Status:      model.Pending,
		CreatedAt:   time.Now(),
		OrderBy:     cli.Token,
	}

	if err := os.orderRepo.Insert(o); err != nil {
		return err
	}

	return nil
}
