package updatestrategy

type defaultSellinUpdateStrategy struct{}

// decrements the sellin by one - default logic
func (d *defaultSellinUpdateStrategy) GetUpdatedSellin(sellIn int) (updatedSellin int) {
	return sellIn - 1
}
