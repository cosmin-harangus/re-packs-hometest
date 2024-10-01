package algo

import (
	"sort"
)

// Calculator algo calculates the minimum number of packs needed to fill an order
// based on the pack sizes provided and the size of the order
//
// It uses a greedy algorithm to calculate the minimum number of packs needed.
// It also allows the user the change the pack sizes dynamically.
type Calculator struct {
	packSizes []int
}

// Shipment contains the pack sizes and the number of packs needed for the order
type Shipment map[int]int

// NewCalculator creates a new calculator with the given pack sizes
func NewCalculator(packSizes []int) *Calculator {
	calc := &Calculator{}
	calc.SetPackSizes(packSizes)
	return calc
}

// GetPackSizes returns the pack sizes for the calculator
func (c *Calculator) GetPackSizes() []int {
	return c.packSizes
}

// SetPacks sets the pack sizes for the calculator
// It sorts the pack sizes in descending order
// and stores them in the calculator
func (c *Calculator) SetPackSizes(packSizes []int) {
	sort.Slice(packSizes, func(i, j int) bool {
		return packSizes[i] > packSizes[j]
	})
	c.packSizes = packSizes
}

// Calculate calculates the minimum number of packs needed to fill an order
func (c *Calculator) Calculate(order int) Shipment {
	n := len(c.packSizes)
	orderSizes := Shipment{}

	// Fill the order with the largest pack sizes first
	for _, size := range c.packSizes {
		if order >= size {
			count := order / size
			order = order % size
			orderSizes[size] = count
		}
	}

	if order > 0 {
		// if there are still items left in the order we need to add a new pack to the order
		// with the smallest pack size that is greater than the order
		doubled := false
		for i := n - 1; i >= 0; i-- {
			size := c.packSizes[i]
			if size > order {
				if _, ok := orderSizes[size]; !ok {
					orderSizes[size] = 1
				} else {
					orderSizes[size] += 1
					doubled = true
				}
				break
			}
		}
		// if we increased a pack we check if we can reduce it by adding a larger pack
		if doubled {
			for i := n - 1; i > 0; i-- {
				hasExactFit := c.packSizes[i-1]%c.packSizes[i] == 0
				exactFit := c.packSizes[i-1] / c.packSizes[i]
				if hasExactFit && orderSizes[c.packSizes[i]] >= exactFit {
					orderSizes[c.packSizes[i-1]] += 1
					orderSizes[c.packSizes[i]] -= exactFit
				}
			}
		}
	}
	// clean up the orderSizes map to remove any pack sizes that are not needed
	for _, size := range c.packSizes {
		if orderSizes[size] == 0 {
			delete(orderSizes, size)
		}
	}

	return orderSizes
}
