package main

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)

	for index, value := range nums {
		compliment := target - value
		if i, ok := indexMap[compliment]; ok {
			return []int{index, i}
		}
		indexMap[value] = index
	}
	return nil
}
