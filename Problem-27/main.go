package main

import "fmt"

func removeElement(nums []int, val int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[cnt] = nums[i]
			cnt++
		}
	}
	fmt.Println(nums[:cnt])
	return cnt
}
