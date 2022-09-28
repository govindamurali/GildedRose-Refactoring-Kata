package gildedrose

import "github.com/emilybache/gildedrose-refactoring-kata/gildedrose/updatestrategy"

func UpdateQuality(items []*Item) {
	itemResolver := updatestrategy.QualityUpdateStrategyResolver{}

	for i := 0; i < len(items); i++ {
		itemUpdateStrategy := itemResolver.GetStrategy(items[i].Name)
		items[0].updateQuality(itemUpdateStrategy)
	}
}

/*
func UpdateQualityBackup(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name != _agedBrie && items[i].Name != _backStage {
			if items[i].Quality > 0 {
				if items[i].Name != _sulfuras {
					items[i].Quality = items[i].Quality - 1
				}
			}
		} else {
			if items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
				if items[i].Name == _backStage {
					if items[i].SellIn < 11 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
					if items[i].SellIn < 6 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
				}
			}
		}

		if items[i].Name != _sulfuras {
			items[i].SellIn = items[i].SellIn - 1
		}

		if items[i].SellIn < 0 {
			if items[i].Name != _agedBrie {
				if items[i].Name != _backStage {
					if items[i].Quality > 0 {
						if items[i].Name != _sulfuras {
							items[i].Quality = items[i].Quality - 1
						}
					}
				} else {
					items[i].Quality = items[i].Quality - items[i].Quality
				}
			} else {
				if items[i].Quality < 50 {
					items[i].Quality = items[i].Quality + 1
				}
			}
		}
	}

}
*/
