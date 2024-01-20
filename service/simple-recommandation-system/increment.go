package simplerecommandationsystem

func (srs *SimpleRecommandationSystem) Increment(menuItemId string, quantity int8) error {
	if err := srs.Sync(); err != nil {
		return err
	}

	srs.Recommands[menuItemId] = srs.Recommands[menuItemId] + int(quantity)

	return nil
}
