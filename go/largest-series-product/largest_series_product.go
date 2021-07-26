package lsproduct

import (
	"errors"
	"strconv"
)

func normalize(in string) ([]int, error) {
	res := make([]int, 0)
	for _, i := range in {
		n, err := strconv.ParseInt(string(i), 10, 32)
		if err != nil {
			return res, err
		}
		res = append(res, int(n))
	}

	return res, nil
}

func product(in []int) int {
	res := 1
	for i := 0; i < len(in); i++ {
		res *= in[i]
	}
	return res
}

func LargestSeriesProduct(in string, span int) (int, error) {
	if span > len(in) || span < 0 {
		return 0, errors.New("span must be smaller than string length")
	}
	if len(in) == 0 {
		return 1, nil
	}
	numbers, err := normalize(in)
	if err != nil {
		return 0, err
	}

	i := 0
	h := 0
	for {
		if i+span > len(in) {
			break
		}
		res := product(numbers[i : i+span])
		if res > h {
			h = res
		}
		i++
	}
	return h, nil
}
