package pangram

import "strings"

func IsPangram(in string) bool {
	letters := make(map[string]int)
	for i := 97; i < 122; i++ {
		letters[string(rune(i))] = 0
	}

	for _, l := range in {
		letters[strings.ToLower(string(l))] = 1
	}

	for _, v := range letters {
		if v == 0 {
			return false
		}
	}
	return true
}
