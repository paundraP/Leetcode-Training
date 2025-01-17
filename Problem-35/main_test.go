package main

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
	}

	for _, test := range tests {
		result := searchInsert(test.nums, test.target)
		if result != test.expected {
			t.Errorf("For input %v and target %d, expected %d but got %d", test.nums, test.target, test.expected, result)
		}
	}
}
