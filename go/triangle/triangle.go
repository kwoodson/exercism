// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
	"sort"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = 0 // not a triangle
	Equ Kind = 1 // equilateral
	Iso Kind = 2 // isosceles
	Sca Kind = 3 // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// valid triangle?
	if a <= 0 || b <= 0 || c <= 0 || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}
	// equilateral check
	if a == b && a == c {
		return Equ
	}
	sides := []float64{a, b, c}
	sort.Slice(sides, func(i, j int) bool {
		return sides[i] < sides[j]
	})
	// isosceles
	if (sides[0] == sides[1] || sides[1] == sides[2]) && sides[0]+sides[1] >= sides[2] {
		return Iso
	}
	// scalene
	if a != b && b != c && a != c && sides[0]+sides[1] >= sides[2] {
		return Sca
	}

	// not a triangle
	return NaT
}
