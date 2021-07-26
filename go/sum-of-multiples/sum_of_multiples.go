package summultiples

func SumMultiples(limit int, d ...int) int {
	if len(d) == 0 {
		return 0
	}
	c := 0
	res := make(map[int]int)
	for _, div := range d {
		if div == 0 {
			continue
		}
		for i := div; i < limit; i += div {
			res[i] = i
		}
	}

	for k := range res {
		c += k
	}
	return c
}
