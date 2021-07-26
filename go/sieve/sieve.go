package sieve

func block(s []int, n int) []int {
	for i := n; i < len(s); i += n {
		s[i] = i
	}
	return s
}

func Sieve(limit int) []int {
	res := make([]int, 0)
	if limit == 1 {
		return res
	}
	s := make([]int, limit+1)
	n := 2
	s[n] = n
	for {
		res = append(res, n)
		s = block(s, n)
		for {
			n++
			if n > limit {
				return res
			}
			if s[n] == 0 {
				break
			}
		}
	}
}
