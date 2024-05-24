package main

func countSubstrings(s string) int {
	count := len(s)

	countPals := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			count++
			left--
			right--
		}
	}
	for i := 0; i < len(s); i++ {
		// Check for odd-length palindromes
		countPals(i, i)

		// Check for even-length palindromes
		if i+1 < len(s) {
			countPals(i, i+1)
		}

	}

	return count
}
