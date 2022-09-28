package updatestrategy

// enums to define the item types
const (
	_agedBrie  string = "Aged Brie"
	_backStage string = "Backstage passes to a TAFKAL80ETC concert"
	_sulfuras  string = "Sulfuras, Hand of Ragnaros"
	_conjured  string = "Conjured Mana Cake"
)

// this
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
	case _conjured:
		return &conjuredUpdateStrategy{}
	}

	return &defaultUpdateStrategy{}
}
