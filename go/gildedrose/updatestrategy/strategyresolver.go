package updatestrategy

type IUpdateStrategy interface {
	GetUpdatedQuality(quality, sellIn int) (updatedQuality int)
	GetUpdatedSellin(sellIn int) (updatedSellin int)
}

type UpdateStrategyResolver struct{}

func (i *UpdateStrategyResolver) GetStrategy(name string) IUpdateStrategy {
	switch name {
	case _agedBrie:
		return &agedBrieUpdateStrategy{}
	case _sulfuras:
		return &sulfurasUpdateStrategy{}
	case _backStage:
		return &backStageUpdateStrategy{}
	}

	return &defaultUpdateStrategy{}
}
