package raindrops

import (
	"strconv"
	"strings"
)

func Convert(drop int) string {
	result := make([]string, 0)
	if drop%3 == 0 {
		result = append(result, "Pling")
	}
	if drop%5 == 0 {
		result = append(result, "Plang")
	}
	if drop%7 == 0 {
		result = append(result, "Plong")
	}
	if len(result) == 0 {
		i := strconv.Itoa(drop)
		result = append(result, i)
	}

	return strings.Join(result, "")
}
