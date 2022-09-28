package updatestrategy

type conjuredUpdateStrategy struct {
	defaultSellinUpdateStrategy // doing this assigns a default sellin update strategy
}

func (a *conjuredUpdateStrategy) GetUpdatedQuality(quality, sellIn int) (updatedQuality int) {
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
