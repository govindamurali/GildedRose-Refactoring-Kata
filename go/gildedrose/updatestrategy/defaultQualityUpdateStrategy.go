package updatestrategy

type defaultQualityUpdateStrategy struct{}

// this is the default quality update strategy mentioned in the doc. will be followed unless we specifically implement for a type
func (a *defaultQualityUpdateStrategy) GetUpdatedQuality(quality, sellIn int) (updatedQuality int) {
	qualityDecrement := -1

	if sellIn < 0 {
		qualityDecrement = -2
	}

	quality = quality + qualityDecrement

	if quality < 0 {
		return 0
	}
	return quality
}
