package updatestrategy

type BackStageUpdateStrategy struct{}

func (a *BackStageUpdateStrategy) GetUpdatedQuality(quality, sellIn int) (updatedQuality int) {
	increment := 1
	if sellIn < 6 {
		increment = 2
	} else if sellIn < 11 {
		increment = 3
	}
	quality = quality + increment
	return quality
}
