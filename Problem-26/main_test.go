package main

import "testing"

func TestRemoveDuplication(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, test := range tests {
		result := removeDuplicates(test.nums)
		if result != test.expected {
			t.Errorf("input: %v\noutput: %v\nexpected: %v", test.nums, result, test.expected)
		}
	}
}
