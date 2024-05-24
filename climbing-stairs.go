package main

func climbStairs(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	a, b := 0, 1

	for i := 0; i < n; i++ {
		temp := b
		b += a
		a = temp
	}
	return b
}
