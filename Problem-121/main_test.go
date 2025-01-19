package main

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
	}

	for _, test := range tests {
		result := maxProfit(test.prices)
		fmt.Printf("result: %d\n", result)
	}
}
