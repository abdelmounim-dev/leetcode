package main

func makeRuneCountMap(s string) map[rune]int {
	runeCountMap := make(map[rune]int)
	for _, v := range s {
		runeCountMap[v]++
	}
	return runeCountMap
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := makeRuneCountMap(s)
	for _, v := range t {
		if _, ok := sMap[v]; !ok {
			return false
		}
		sMap[v]--
		if sMap[v] < 0 {
			return false
		}
	}

	return true
}
