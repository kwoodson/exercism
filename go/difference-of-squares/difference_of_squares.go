package diffsquares

import "math"

func SquareOfSum(n int) int {
	d := make(map[int]int)

	d[0] = 0
	d[1] = 0
	res := 0
	i := 0
	for i = 1; i < n; i++ {
		d[i] += i + d[i-1]
		res += d[i]
	}
	// 100 want 225
	return int(math.Pow(float64(d[i-1]+n), 2))

}

func SumOfSquares(n int) int {
	d := make(map[int]int)

	d[0] = 0
	d[1] = 1
	res := 0
	i := 0
	for i = 2; i < n+1; i++ {
		d[i] += int(math.Pow(float64(i), 2)) + d[i-1]
		res += d[i]
	}

	return d[i-1]
}

func Difference(a int) int {
	return SquareOfSum(a) - SumOfSquares(a)
}
