// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	// d := time.Millisecond * 1000 * 1000000000
	t = t.Add(time.Millisecond * 1000 * 1000000000)
	return t
}