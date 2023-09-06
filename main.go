package main

import (
	itemfilter "filter/itemFilter"
	"fmt"
)

func main() {
	items := []itemfilter.Item{
		{"Item1", 10.99, 5},
		{"Item2", 5.99, 10},
		{"Item3", 4.50, 3},
		{"Item4", 2.50, 2},
		{"Item5", 6.50, 12},
		{"Item6", 10.50, 0},
		{"Item7", 14.20, 3},
		{"Item8", 12.50, 1},
		{"Item9", 14.50, 6},
		{"Item10", 20.50, 3},

		// Add more items as needed
	}

	options := itemfilter.FilterOptions{
		MaxPrice:    15.0,
		MinQuantity: 5,
	}

	filteredItems, err := itemfilter.Filter(items, options)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Filtered Items:")
	itemfilter.PrintItems(filteredItems)
}
