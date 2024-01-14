package simplerecommandationsystem

func (srs *SimpleRecommandationSystem) Init() error {
	menus, err := srs.MenuRepo.GetAll()
	if err != nil {
		return err
	}

	for _, menu := range menus {
		srs.Recommands[menu.Id] = 0
	}

	return nil
}
