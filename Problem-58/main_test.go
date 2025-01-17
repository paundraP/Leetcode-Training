package main

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
		{"", 0},
		{"Hello", 5},
		{"   ", 0},
	}

	for _, test := range tests {
		result := lengthOfLastWord(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %d but got %d", test.input, test.expected, result)
		}
	}
}
