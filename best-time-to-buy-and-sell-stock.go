package main

func maxProfit(prices []int) int {
	var maxDiff struct {
		min  int
		max  int
		diff int
	}
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > maxDiff.diff {
				maxDiff.diff = prices[j] - prices[i]
				maxDiff.min = i
				maxDiff.max = j
			}
		}
	}

	return maxDiff.diff
}

func maxProfitOneLoop(prices []int) int {
	var maxDiff struct {
		min  int
		max  int
		diff int
	}
	maxDiff.min = prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < maxDiff.min {
			maxDiff.min = prices[i]
			maxDiff.max = 0
		} else if prices[i] > maxDiff.max {
			maxDiff.max = prices[i]
			if prices[i]-maxDiff.min > maxDiff.diff {
				maxDiff.diff = prices[i] - maxDiff.min
			}
		}
	}

	return maxDiff.diff
}
