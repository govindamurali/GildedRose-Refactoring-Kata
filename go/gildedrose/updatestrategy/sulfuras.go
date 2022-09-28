package updatestrategy

type sulfurasUpdateStrategy struct{}

func (a *sulfurasUpdateStrategy) GetUpdatedQuality(quality, sellIn int) (updatedQuality int) {
	return quality
}
func (d *sulfurasUpdateStrategy) GetUpdatedSellin(sellIn int) (updatedSellin int) {
	return sellIn
}
