package gildedrose

import "github.com/emilybache/gildedrose-refactoring-kata/gildedrose/updatestrategy"

// primary method to update quality
// @param items - takes in a list of items to update
func UpdateQuality(items []*Item) {
	updateStrategyResolver := updatestrategy.UpdateStrategyResolver{}

	for i := 0; i < len(items); i++ {
		strategy := updateStrategyResolver.GetStrategy(items[i].Name)
		items[i].updateSellin(strategy)
		items[i].updateQuality(strategy)
	}
}
