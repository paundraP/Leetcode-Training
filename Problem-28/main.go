package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		match := true

		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}

func main() {
	test1 := strStr("leetcode", "leeto")
	test2 := strStr("sadbutsad", "sad")
	fmt.Println(test1)
	fmt.Println(test2)
}
