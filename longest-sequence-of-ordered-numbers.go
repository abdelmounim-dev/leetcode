package main

type sequence struct {
	start int
	end   int
}

func (s *sequence) len() int {
	return s.end - s.start
}

func longestSequence(nums []int) int {
	longestSequence := sequence{0, 0}
	currentSequence := sequence{0, 0}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if longestSequence.len() < currentSequence.len() {
				longestSequence = currentSequence
			}
			currentSequence = sequence{i, i}
		}

		if nums[i] < nums[i+1] {
			currentSequence.end = i
		}
	}

	return longestSequence.end - longestSequence.start
}
