package main

func containsDuplicate(nums []int) bool {
	elements := make(map[int]int, len(nums))

	for _, v := range nums {
		if _, ok := elements[v]; ok {
			return true
		}
		elements[v] = 1
	}

	return false
}
