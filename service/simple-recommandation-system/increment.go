package simplerecommandationsystem

func (srs *SimpleRecommandationSystem) Increment(menuItemId string, quantity int8) {
	srs.Recommands[menuItemId] = srs.Recommands[menuItemId] + int(quantity)
}
