package main

import (
	"fmt"
)

// Item represents a product in the shopping cart
type Item struct {
	Name  string
	Price float64
	Qty   int
}

func main() {
	// Shopping cart represented as a map with item names as keys
	cart := map[string]Item{
		"apple":  {Name: "Apple", Price: 0.5, Qty: 4},
		"banana": {Name: "Banana", Price: 0.3, Qty: 6},
		"milk":   {Name: "Milk", Price: 1.2, Qty: 2},
	}

	totalCost := 0.0

	fmt.Println("Shopping Cart:")
	for _, item := range cart {
		itemTotal := item.Price * float64(item.Qty)
		totalCost += itemTotal
		fmt.Printf("%s - %d pcs - $%.2f each - Total: $%.2f\n", item.Name, item.Qty, item.Price, itemTotal)
	}

	fmt.Printf("Total Cost: $%.2f\n", totalCost)
}

