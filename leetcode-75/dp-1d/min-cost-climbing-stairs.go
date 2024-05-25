package main

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}
	if len(cost) == 1 {
		return cost[0]
	}
	if len(cost) == 2 {
		return cost[1]
	}
	c1, c2, c3 := 0, cost[0], cost[1]
	for i := 3; i <= len(cost); i++ {
		c1 = c2
		c2 = c3
		c3 = cost[i] + min(c1, c2)
	}
	return c3
}
