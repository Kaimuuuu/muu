package simplerecommandationsystem

func (srs *SimpleRecommandationSystem) Sync() error {
	menuItems, err := srs.MenuRepo.All()
	if err != nil {
		return err
	}

	for _, m := range menuItems {
		_, ok := srs.Recommands[m.Id]
		if !ok {
			srs.Recommands[m.Id] = 0
		}
	}

	return nil
}
