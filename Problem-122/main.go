package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min, max, tmp := prices[0], 0, 0

	for _, price := range prices {
		if price < min {
			min = price
		}
		if price-min > tmp {
			tmp = price - min
			max += tmp
			tmp = 0
			min = price
		}
	}
	return max
}
