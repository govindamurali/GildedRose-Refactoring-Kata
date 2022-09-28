package updatestrategy

// this update strategy will be used for unknown types, quality update logic & sellin update logic both are default
type defaultUpdateStrategy struct {
	defaultSellinUpdateStrategy
	defaultQualityUpdateStrategy
}

type defaultSellinUpdateStrategy struct{}

// decrements the sellin by one - default logic
func (d *defaultSellinUpdateStrategy) GetUpdatedSellin(sellIn int) (updatedSellin int) {
	return sellIn - 1
}

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
