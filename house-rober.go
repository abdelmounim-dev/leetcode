package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	if len(nums) == 3 {
		return max(nums[0]+nums[2], nums[1])
	}
	h0, h1, h2, h3 := 0, nums[0], nums[1], nums[0]+nums[2]

	for i := 3; i < len(nums); i++ {
		h0 = h1
		h1 = h2
		h2 = h3
		h3 = nums[i] + max(h1, h0)
	}

	return max(h2, h3)
}

