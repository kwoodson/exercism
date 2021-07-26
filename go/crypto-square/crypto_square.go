package cryptosquare

import (
	"strings"
)

func FormRectangle(s int) (int, int) {
	for r := 1; r < s; r++ {
		for c := r; c < (r + 2); c++ {
			if c*r >= s {
				if c >= r && c-r <= 1 {
					return c, r
				}
			}
		}
	}

	return 0, 0
}

func Shift(in string, l, h int) string {
	inTmp := make([]string, h)
	c := 0
	for i := 0; i < l; i++ {
		for y := 0; y < h; y++ {
			if c > len(in)-1 {
				inTmp[y] += " "
			} else {
				inTmp[y] += string(in[c])
			}
			c += 1
		}
	}

	return strings.Join(inTmp, " ")
}

func Encode(in string) string {
	if len(in) == 1 {
		return in
	}

	in = strings.ToLower(in)
	res := ""
	for _, z := range strings.ToLower(in) {
		c := int(z)
		if c >= 97 && c <= 122 || c >= 49 && c <= 57 {
			res += string(z)
		}
	}
	h, l := FormRectangle(len(res))
	r := Shift(res, l, h)

	return r
}
