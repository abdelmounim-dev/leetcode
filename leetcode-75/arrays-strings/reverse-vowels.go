package main

func reverseVowels(s string) string {
	vowels := map[rune]int8{
		'a': 0, 'e': 0, 'o': 0, 'u': 0, 'i': 0,
	}
	stack := make([]rune, len(s))
	for _, v := range s {
		if _, ok := vowels[v]; ok {
			stack = append(stack, v)
		}
	}
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		if _, ok := vowels[runes[i]]; ok {
			runes[i] = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}

	return string(runes)
}
