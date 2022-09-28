package updatestrategy

type sulfurasUpdateStrategy struct{}

func (a *sulfurasUpdateStrategy) GetUpdatedQuality(quality, sellIn int) (updatedQuality int) {
	return quality
}

// only in case of sulfuras, we have to specifically define a separate update logic, hence we implement it specifically
func (d *sulfurasUpdateStrategy) GetUpdatedSellin(sellIn int) (updatedSellin int) {
	return sellIn
}
