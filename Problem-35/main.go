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
	array := []int{1, 3, 5, 6}
	test1 := searchInsert(array, 5)
	test2 := searchInsert(array, 2)
	test3 := searchInsert(array, 7)
	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
}
