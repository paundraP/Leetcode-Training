package main

import "fmt"

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	pos := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && pos > 0 {
			return pos
		} else if s[i] != ' ' {
			pos++
		}
	}
	return pos
}

func main() {
	fmt.Println("Hello World")
}
