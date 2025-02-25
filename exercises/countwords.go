package exercises

import "strings"

func CountWords(text, target_word string) int {
	if len(target_word) == 0 {
		return 0
	}

	if len(text) == 0 {
		return 0
	}

	counter := map[string]int{}
	words := strings.Fields(text)

	for _, word := range words {
		counter[strings.ToLower(word)]++
	}

	return counter[target_word]
}