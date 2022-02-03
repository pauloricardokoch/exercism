package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}

	second := false
	sum := 0

	for i := len(id) - 1; i > -1; i-- {
		value, err := strconv.Atoi(id[i : i+1])
		if err != nil {
			return false
		}

		if second {
			value = value * 2
			if value > 9 {
				value -= 9
			}
		}

		sum += value
		second = !second
	}

	return sum%10 == 0
}
