package main

func countSubstrings(s string) int {
	count := len(s)
	for i := 0; i < len(s); i++ {
		// Check for odd-length palindromes
		var j int
		for j = 0; i-j >= 0 && i+j < len(s) && s[i-j] == s[i+j]; j++ {
		}
		if 2*j-1 > 1 {
			count += j - 1
		}

		// Check for even-length palindromes
		if i+1 < len(s) && s[i+1] == s[i] {
			for j = 0; i-j >= 0 && i+j+1 < len(s) && s[i-j] == s[i+j+1]; j++ {
			}
			if 2*j > 1 {
				count += j
			}
		}
	}

	return count
}
