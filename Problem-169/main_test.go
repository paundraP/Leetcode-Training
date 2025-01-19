package main

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 2, 2}, 2},
	}
	for _, test := range tests {
		result := majorityElement(test.nums)
		if result != test.expected {
			t.Errorf("input: %d\noutput: %d\nexpected: %d", test.nums, result, test.expected)
		}
	}
}
