package simplerecommandationsystem

func (srs *SimpleRecommandationSystem) Sync() error {
	menus, err := srs.MenuRepo.GetAll()
	if err != nil {
		return err
	}

	for _, menu := range menus {
		_, ok := srs.Recommands[menu.Id]
		if !ok {
			srs.Recommands[menu.Id] = 0
		}
	}

	return nil
}
