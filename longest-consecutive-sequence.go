package main

func makeMapFromArray(arr []int) map[int]int {
	m := make(map[int]int, len(arr))
	for _, val := range arr {
		m[val]++
	}
	return m
}

func longestConsecutive(nums []int) int {
	numsMap := makeMapFromArray(nums)
	seqLen := 1
	for _, v := range nums {
		curLen := 0
		if _, ok := numsMap[v]; ok {
			for w := v; ok; w++ {
				delete(numsMap, w)
				curLen++
				_, ok = numsMap[w]
			}
		}
		if curLen > seqLen {
			seqLen = curLen
		}
	}

	return seqLen
}
