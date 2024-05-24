package main

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n < 3 {
		return 1
	}
	v0, v1, v2, v3 := 0, 1, 1, 2
	for i := 4; i < n+1; i++ {
		v0 = v1
		v1 = v2
		v2 = v3
		v3 = v0 + v1 + v2
	}
	return v3
}
