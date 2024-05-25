package main

import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	threshold := slices.Max(candies)
	isKWGNC := make([]bool, len(candies))
	for i, c := range candies {
		isKWGNC[i] = c+extraCandies >= threshold
	}

	return isKWGNC
}
