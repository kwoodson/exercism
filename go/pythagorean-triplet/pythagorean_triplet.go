package pythagorean

import "math"

type Triplet [3]int

func Range(min, max int) []Triplet {
	t := make([]Triplet, 0)
	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			for c := b + 1; c <= max; c++ {
				if math.Pow(float64(a), 2)+math.Pow(float64(b), 2) == math.Pow(float64(c), 2) {
					t = append(t, Triplet{a, b, c})
				}
			}
		}
	}

	return t
}

func Sum(p int) []Triplet {
	res := make([]Triplet, 0)
	triplets := Range(1, p/2)
	for _, t := range triplets {
		if t[0]+t[1]+t[2] == p {
			res = append(res, t)
		}
	}

	return res
}
