package gildedrose

import "github.com/emilybache/gildedrose-refactoring-kata/gildedrose/updatestrategy"

type Item struct {
	Name            string
	SellIn, Quality int
}

func (i *Item) updateQuality(itemUpdateStrategy updatestrategy.IUpdateStrategy) {
	i.Quality = itemUpdateStrategy.GetUpdatedQuality(i.Quality, i.SellIn)
}
func (i *Item) updateSellin(itemUpdateStrategy updatestrategy.IUpdateStrategy) {
	i.SellIn = itemUpdateStrategy.GetUpdatedSellin(i.SellIn)
}
