package luhn

import (
	"strconv"
	"strings"
)

func convertToNumber(n string) ([]int, error) {
	num := make([]int, 0)
	n = strings.Trim(n, " -")
	n = strings.ReplaceAll(n, " ", "")
	for _, c := range n {
		// convert c to integer if possible
		z, err := strconv.ParseInt(string(c), 10, 32)
		if err != nil {
			return num, err
		}
		num = append(num, int(z))
	}
	return num, nil
}

func Valid(cc string) bool {
	num, err := convertToNumber(cc)
	if err != nil || len(num) < 2 {
		return false
	}

	for i := len(num) - 2; i >= 0; i -= 2 {
		z := num[i] * 2
		if z > 9 {
			z -= 9
		}
		num[i] = z
	}

	sum := 0
	for _, n := range num {
		sum += n
	}
	if sum%10 == 0 {
		return true
	}

	return false
}
