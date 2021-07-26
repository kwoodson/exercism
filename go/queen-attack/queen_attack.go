package queenattack

import (
	"errors"
	"strconv"
)

type position struct {
	rank int // row
	file int // column
}

func ParsePosition(in string) (position, error) {
	p := position{}
	if len(in) != 2 {
		return p, errors.New("invalid position")
	}
	// verify first letter is valid
	f := rune(in[0])
	file := int(f) - 96
	if file < 1 || file > 8 {
		return p, errors.New("invalid file inside of position")
	}
	rank, err := strconv.ParseInt(in[1:], 10, 32)
	if err != nil {
		return p, errors.New("invalid integer inside of position location")
	}
	if rank < 1 || rank > 8 {
		return p, errors.New("invalid rank inside of position")
	}
	p.rank = int(rank)
	p.file = file
	return p, nil
}

func checkDiagonal(wp, p position) bool {
	x, y := wp.rank, wp.file
	for {
		x--
		y--
		if x < 1 || y < 1 {
			break
		}
		if x == p.rank && y == p.file {
			return true
		}
	}

	x, y = wp.rank, wp.file
	for {
		x++
		y++
		if x > 8 || y > 8 {
			break
		}
		if x == p.rank && y == p.file {
			return true
		}
	}

	x, y = wp.rank, wp.file
	for {
		x--
		y++
		if x < 1 || y > 8 {
			break
		}
		if x == p.rank && y == p.file {
			return true
		}
	}
	x, y = wp.rank, wp.file
	for {
		x++
		y--
		if x > 8 || y < 1 {
			break
		}
		if x == p.rank && y == p.file {
			return true
		}
	}

	return false
}

func CanQueenAttack(w, b string) (bool, error) {
	wp, err := ParsePosition(w)
	if err != nil {
		return false, err
	}

	bp, err := ParsePosition(b)
	if err != nil {
		return false, err
	}

	if wp.file == bp.file && wp.rank == bp.rank {
		return false, errors.New("pieces located on same square")
	}

	if wp.file == bp.file || wp.rank == bp.rank {
		return true, nil
	}

	// determine diagonal
	return checkDiagonal(wp, bp), nil
}
