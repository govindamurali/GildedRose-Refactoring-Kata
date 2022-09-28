package updatestrategy

type sulfurasUpdateStrategy struct{}

func (a *sulfurasUpdateStrategy) GetUpdatedQuality(quality, sellIn int) (updatedQuality int) {
	return quality
}
