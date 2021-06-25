package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (Matrix, error) {
	m := make(Matrix, 0)
	for i, line := range strings.Split(s, "\n") {
		row := make([]int, 0)
		for _, num := range strings.Fields(line) {
			n, err := strconv.ParseInt(num, 10, 32)
			if err != nil {
				return m, err
			}
			row = append(row, int(n))
		}
		// validate rows are same size
		if i > 0 && len(m[0]) != len(row) {
			return m, errors.New("unequal row length")
		}
		m = append(m, row)
	}
	return m, nil
}

func (m Matrix) Rows() [][]int {
	rows := make([][]int, 0)
	for _, r := range m {
		row := make([]int, 0)
		row = append(row, r...)
		rows = append(rows, row)
	}
	return rows
}

func (m Matrix) Cols() [][]int {
	cols := make([][]int, 0)
	for r := 0; r < len(m[0]); r++ {
		col := make([]int, 0)
		for c := 0; c < len(m); c++ {
			col = append(col, m[c][r])
		}
		cols = append(cols, col)
	}
	return cols
}

func (m Matrix) Set(r, c, val int) bool {
	if r < 0 || r > len(m)-1 || c < 0 || c > len(m[r])-1 {
		return false
	}
	m[r][c] = val
	return true
}
