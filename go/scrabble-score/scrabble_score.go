package scrabble

import "strings"

var lv = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

func Score(word string) int {
	score := 0
	for _, char := range strings.ToUpper(word) {
		for letter, value := range lv {
			if strings.ContainsRune(letter, char) {
				score += value
			}
		}
	}

	return score
}
