package updatestrategy

type backStageUpdateStrategy struct {
	defaultSellinUpdateStrategy // doing this assigns a default sellin update strategy
}

func (a *backStageUpdateStrategy) GetUpdatedQuality(quality, sellIn int) (updatedQuality int) {
	if sellIn < 0 {
		return 0
	}

	increment := 1
	if sellIn < 6 {
		increment += 1
	}
	if sellIn < 11 {
		increment += 1
	}
	quality = quality + increment

	if quality >= 50 {
		quality = 50
	}

	return quality
}
