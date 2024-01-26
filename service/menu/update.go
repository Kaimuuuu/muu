package menu

import (
	"fmt"
	"os"
)

func (ms *MenuService) Update(menuItemId string, req UpdateMenuRequest) error {
	m, err := ms.menuRepo.GetById(menuItemId)
	if err != nil {
		return err
	}

	os.Remove(fmt.Sprintf("public/%s", m.ImagePath))

	m.Catagory = req.Catagory
	m.Description = req.Description
	m.Name = req.Name
	m.Price = req.Price
	m.ImagePath = req.ImagePath

	if err := ms.menuRepo.Update(menuItemId, m); err != nil {
		return err
	}

	return nil
}
