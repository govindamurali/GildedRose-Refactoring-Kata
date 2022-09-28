package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
	"github.com/stretchr/testify/assert"
)

var items = []*gildedrose.Item{
	{"+5 Dexterity Vest", 10, 20},
	{"Aged Brie", 2, 0},
	{"Elixir of the Mongoose", 5, 7},
	{"Sulfuras, Hand of Ragnaros", 0, 80},
	{"Sulfuras, Hand of Ragnaros", -1, 80},
	{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
	{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
	{"Conjured Mana Cake", 3, 6}, // <-- :O
}

var updatedItems = []*gildedrose.Item{
	{"+5 Dexterity Vest", 9, 19},
	{"Aged Brie", 1, 1},
	{"Elixir of the Mongoose", 4, 6},
	{"Sulfuras, Hand of Ragnaros", 0, 80},
	{"Sulfuras, Hand of Ragnaros", -1, 80},
	{"Backstage passes to a TAFKAL80ETC concert", 14, 21},
	{"Backstage passes to a TAFKAL80ETC concert", 9, 50},
	{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
	{"Conjured Mana Cake", 2, 4}, // <-- :O
}

func TestUpdateQuality(t *testing.T) {
	gildedrose.UpdateQuality(items)

	for key := range items {
		assert.Equal(t, updatedItems[key].Quality, items[key].Quality)
		assert.Equal(t, updatedItems[key].SellIn, items[key].SellIn)
		assert.Equal(t, updatedItems[key].Name, items[key].Name)
	}
}
