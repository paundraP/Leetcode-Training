package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min, max := prices[0], 0

	for _, price := range prices {
		if price < min {
			min = price
		} else if price-min > max {
			max = price - min
		}
	}
	return max
}
