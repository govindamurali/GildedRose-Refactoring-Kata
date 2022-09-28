package gildedrose

import "github.com/emilybache/gildedrose-refactoring-kata/gildedrose/updatestrategy"

type Item struct {
	Name            string
	SellIn, Quality int
}

// updates the quality of the item using a particular updateStrategy that's passed
func (i *Item) updateQuality(itemUpdateStrategy updatestrategy.IUpdateStrategy) {
	i.Quality = itemUpdateStrategy.GetUpdatedQuality(i.Quality, i.SellIn)
}

// updates the sellin of the item using a particular updateStrategy that's passed
func (i *Item) updateSellin(itemUpdateStrategy updatestrategy.IUpdateStrategy) {
	i.SellIn = itemUpdateStrategy.GetUpdatedSellin(i.SellIn)
}
