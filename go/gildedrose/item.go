package gildedrose

import "github.com/emilybache/gildedrose-refactoring-kata/gildedrose/updatestrategy"

type Item struct {
	Name            string
	SellIn, Quality int
}

// default implementation of updateQuality.
func (i *Item) updateQuality(itemUpdateStrategy updatestrategy.IUpdateStrategy) {
	i.Quality = itemUpdateStrategy.GetUpdatedQuality(i.Quality, i.SellIn)
}
