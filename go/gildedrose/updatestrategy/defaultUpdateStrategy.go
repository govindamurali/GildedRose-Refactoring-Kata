package updatestrategy

type DefaultUpdateStrategy struct{}

func (a *DefaultUpdateStrategy) GetUpdatedQuality(quality, sellIn int) (updatedQuality int) {
	qualityDecrement := 1

	if sellIn < 0 {
		qualityDecrement = 2
	}

	quality = quality + qualityDecrement

	if quality < 0 {
		return 0
	}
	return quality
}