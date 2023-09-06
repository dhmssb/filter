// itemfilter.go
package itemfilter

import (
	"errors"
	"fmt"
)

// Item represents an item with a name, price, and quantity.
type Item struct {
	Name     string
	Price    float64
	Quantity int
}

// FilterOptions contains filtering criteria.
type FilterOptions struct {
	MaxPrice    float64
	MinQuantity int
}

// Filter filters items based on the provided options.
func Filter(items []Item, options FilterOptions) ([]Item, error) {
	if options.MaxPrice < 0 || options.MinQuantity < 0 {
		return nil, errors.New("MaxPrice and MinQuantity must be non-negative")
	}

	filteredItems := []Item{}
	for _, item := range items {
		if item.Price <= options.MaxPrice && item.Quantity >= options.MinQuantity {
			filteredItems = append(filteredItems, item)
		}
	}
	return filteredItems, nil
}

// PrintItems prints the filtered items to the console.
func PrintItems(filteredItems []Item) {
	for _, item := range filteredItems {
		fmt.Printf("Name: %s, Price: %.2f, Quantity: %d\n", item.Name, item.Price, item.Quantity)
	}
}
