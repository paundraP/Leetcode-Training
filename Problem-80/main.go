package main

import "fmt"

func removeDuplicates(nums []int) int {
	j := 1
	cnt := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			cnt++
		} else {
			cnt = 1
		}
		if cnt <= 2 {
			nums[j] = nums[i]
			j++
		}
	}
	fmt.Println(nums[:j])
	return j
}
