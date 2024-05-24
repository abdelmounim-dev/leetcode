package main

func longestPalindrome(s string) string {
	longestPal := ""
	for i := 0; i < len(s); i++ {
		var j int
		for j = i; j < len(s)-i-1 && i-j > 0 && s[i-j] == s[i+j]; j++ {
		}
		if j*2 > len(longestPal) {
			longestPal = s[i-j : i+j]
		}
	}
	return longestPal
}
