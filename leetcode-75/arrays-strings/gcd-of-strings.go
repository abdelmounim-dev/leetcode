package main

func isDivisor(str string, div string) bool {
	if len(div) == 0 || len(str)%len(div) > 0 {
		return false
	}
	for i := 0; i < len(str); i += len(div) {
		if str[i:i+len(div)] != div {
			return false
		}
	}
	return true
}

func gcdOfStrings(str1 string, str2 string) string {
	minLen := min(len(str1), len(str2))
	for i := minLen; i > 0; i-- {
		div := str1[:i]
		if isDivisor(str1, div) && isDivisor(str2, div) {
			return div
		}
	}

	return ""
}
