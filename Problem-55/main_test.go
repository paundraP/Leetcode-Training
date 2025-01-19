package main

import (
	"testing"
)

func TestCanJump(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, true},
	}
	for _, test := range tests {
		result := canJump(test.nums)
		if result != test.expected {
			t.Errorf("yahh gagal")
		}
	}
}
