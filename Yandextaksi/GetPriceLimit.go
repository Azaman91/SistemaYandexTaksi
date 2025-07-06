package main

const(
	minPrice = 99.00
	maxPrice = 20000
)

func ApplyPriceLimits(price float64)float64{
	if price < minPrice{
		return minPrice
	}
	if price > maxPrice{
		return maxPrice
	}
	return price
}
