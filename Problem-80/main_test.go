package main

import "testing"

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 5},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
	}
	for _, test := range tests {
		result := removeDuplicates(test.nums)
		if result != test.expected {
			t.Errorf("nums: %v \nresult: %d \nexpected: %d\n", test.nums, result, test.expected)
		}
	}
}
