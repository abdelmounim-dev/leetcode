package main

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	start := 0
	chars := make(map[byte]int, len(s))
	for end := 0; end < len(s); end++ {
		if index, found := chars[s[end]]; found && index >= start {
			start = index + 1
		}
		chars[s[end]] = end
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}

	}
	return maxLength
}
