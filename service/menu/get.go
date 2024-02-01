package menu

import (
	"kaimuu/model"
	"kaimuu/service/promotion"
)

func (ms *MenuService) Get(cli *model.Client) ([]promotion.PromotionMenuItemResponse, error) {
	menu, err := ms.promotionServ.GetMenu(cli.PromotionId)
	if err != nil {
		return []promotion.PromotionMenuItemResponse{}, err
	}

	return menu, nil
}
