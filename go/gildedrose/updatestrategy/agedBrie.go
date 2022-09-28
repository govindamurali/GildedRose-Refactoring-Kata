package updatestrategy

type AgedBrieUpdateStrategy struct{}

func (a *AgedBrieUpdateStrategy) GetUpdatedQuality(quality, sellIn int) int {
	quality = quality + 1
	if quality >= 50 {
		quality = 50
	}
	return quality
}
