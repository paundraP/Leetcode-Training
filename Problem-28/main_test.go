package main

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		expected int
	}{
		{"leetcode", "leeto", -1},
		{"sadbudsad", "sad", 0},
	}
	for _, test := range tests {
		result := strStr(test.haystack, test.needle)
		if result != test.expected {
			t.Errorf("For input %s and target %s, expected %d but got %d", test.haystack, test.needle, test.expected, result)
		}
	}
}
