package main

import "testing"

func TestJump(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
	}
	for _, test := range tests {
		result := jump(test.nums)
		if result != test.expected {
			t.Errorf("error at test with input: %v, expected: %d, got: %d", test.nums, test.expected, result)
		}
	}
}
