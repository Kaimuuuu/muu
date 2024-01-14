package menu

import (
	"kaimuu/model"
	"time"

	"github.com/google/uuid"
)

func (ms *MenuService) CreateMenu(req CreateMenuRequest, employeeId string) error {
	m := &model.MenuItem{
		Id:          uuid.NewString(),
		Name:        req.Name,
		Description: req.Description,
		Catagory:    req.Catagory,
		Weight:      req.Weight,
		Price:       req.Price,
		OutOfStock:  req.OutOfStock,
		ImagePath:   req.ImagePath,
		CreatedAt:   time.Now(),
		CreatedBy:   employeeId,
	}

	if err := ms.menuRepo.Insert(m); err != nil {
		return err
	}

	return nil
}
