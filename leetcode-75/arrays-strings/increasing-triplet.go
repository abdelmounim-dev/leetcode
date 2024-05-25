package main

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	n1, n2 := nums[0], int(^uint(0)>>1)
	for i := 1; i < len(nums); i++ {
		if nums[i] <= n1 {
			n1 = nums[i]
			continue
		}
		if nums[i] > n2 {
			return true
		}
		if nums[i] < n2 {
			n2 = nums[i]
		}
	}
	return false
}

