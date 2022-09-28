package updatestrategy

type defaultSellinUpdateStrategy struct {
}

func (d *defaultSellinUpdateStrategy) GetUpdatedSellin(sellIn int) (updatedSellin int) {
	return sellIn - 1
}
