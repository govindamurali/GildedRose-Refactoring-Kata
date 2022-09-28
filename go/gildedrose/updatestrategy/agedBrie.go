package updatestrategy

type agedBrieUpdateStrategy struct {
	defaultSellinUpdateStrategy
}

func (a *agedBrieUpdateStrategy) GetUpdatedQuality(quality, sellIn int) int {
	quality = quality + 1
	if quality >= 50 {
		quality = 50
	}
	return quality
}
