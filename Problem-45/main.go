package main

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	steps := 0
	far := 0
	curr_step := 0

	for i := 0; i < len(nums)-1; i++ {
		if far < i+nums[i] {
			far = i + nums[i]
		}
		if i == curr_step {
			steps++
			curr_step = far
			if curr_step >= len(nums)-1 {
				break
			}
		}
	}
	return steps
}
