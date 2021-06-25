package isogram

import "strings"

func IsIsogram(s string) bool {
	// if s == "" {
	// 	return true
	// }
	d := make(map[string]int)
	for _, l := range s {
		l := strings.ToLower(string(l))
		d[l] += 1
		if l == " " || l == "-" { // can repeat
			continue
		}
		if d[l] > 1 {
			return false
		}
	}

	return true
}
