package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	word = strings.ReplaceAll(word, " ", "")
	word = strings.ReplaceAll(word, "-", "")
	word = strings.ToLower(word)

	for i := 1; i < len(word); i++ {
		letter := word[i : i+1]
		before := word[:i]

		if strings.Contains(before, letter) {
			return false
		}
	}

	return true
}
