package main

const mod = 1000000007

func numTilings(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 5
	}
	t1, t2, t3 := 1, 2, 5

	for i := 4; i <= n; i++ {
		t1, t2, t3 = t2, t3, (2*t3+t1)%mod
	}

	return t3
}

