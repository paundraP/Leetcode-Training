package main

func canJump(nums []int) bool {
	m := 0
	for i := 0; i < len(nums); i++ {
		if i > m {
			return false
		}
		if i+nums[i] > m {
			m = i + nums[i]
		}
		if m >= len(nums)-1 {
			return true
		}
	}
	return false
}
