package main

import "fmt"

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func main() {
	fmt.Println("Hello World")
}
