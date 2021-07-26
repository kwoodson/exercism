package secret

import (
	"strconv"
	"strings"
)

var HandShakes = map[int]string{
	0: "wink",
	1: "double blink",
	2: "close your eyes",
	3: "jump",
}

func Handshake(shake uint) []string {
	s := strconv.FormatInt(int64(shake), 2)
	rev := false
	if len(s) == 5 && strings.HasPrefix(s, "1") {
		s = s[1:]
		rev = true
	}
	result := make([]string, 0)
	pos := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == []byte("1")[0] {
			result = append(result, HandShakes[pos])
		}
		pos++
	}

	if rev {
		r := make([]string, 0)
		for i := len(result) - 1; i >= 0; i-- {
			r = append(r, result[i])
		}
		result = r
	}
	return result
}
