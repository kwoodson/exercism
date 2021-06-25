package grains

import (
	"errors"
)

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("Invalid value for square.")
	}
	d := make([]uint64, 0)
	d = append(d, 0, 1, 2)

	for i := 3; i < n+1; i++ {
		v := d[i-1] * 2
		d = append(d, uint64(v))
	}

	return d[n], nil
}

func Total() uint64 {
	var total uint64 = 0
	for i := 1; i < 65; i++ {
		z, _ := Square(i)
		total += z
	}
	return total
}
