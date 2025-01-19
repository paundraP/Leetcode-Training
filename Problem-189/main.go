package main

import "fmt"

func Rotate(nums []int, k int) {
	k = k % len(nums)

	reverse(nums, 0, len(nums)-1)
	fmt.Printf("first reverse: %d\n", nums)
	reverse(nums, 0, k-1)
	fmt.Printf("second reverse: %d\n", nums)
	reverse(nums, k, len(nums)-1)
	fmt.Printf("third reverse: %d\n", nums)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
