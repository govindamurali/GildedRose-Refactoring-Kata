package updatestrategy

type IUpdateStrategy interface {
	GetUpdatedQuality(quality, sellIn int) (updatedQuality int)
}

type QualityUpdateStrategyResolver struct{}

func (i *QualityUpdateStrategyResolver) GetStrategy(name string) IUpdateStrategy {
	switch name {
	case _agedBrie:
		return &AgedBrieUpdateStrategy{}
	case _sulfuras:
		return &sulfurasUpdateStrategy{}
	case _backStage:
		return &BackStageUpdateStrategy{}
	}

	return &DefaultUpdateStrategy{}
}
