package updatestrategy

type ConjuredUpdateStrategy struct {
	defaultSellinUpdateStrategy
}

func (a *ConjuredUpdateStrategy) GetUpdatedQuality(quality, sellIn int) (updatedQuality int) {
	qualityIncrement := -2

	if sellIn < 0 {
		qualityIncrement = -4
	}

	quality = quality + qualityIncrement

	if quality < 0 {
		return 0
	}
	return quality
}
