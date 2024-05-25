package main

func mergeAlternately(word1 string, word2 string) string {
	mergedWord := []byte{}
	i := 0
	for ; i < len(word1) && i < len(word2); i++ {
		mergedWord = append(mergedWord, word1[i], word2[i])
	}
	if i < len(word1) {
		for ; i < len(word1); i++ {
			mergedWord = append(mergedWord, word1[i])
		}
	}
	if i < len(word2) {
		for ; i < len(word2); i++ {
			mergedWord = append(mergedWord, word2[i])
		}
	}
	return string(mergedWord)
}
