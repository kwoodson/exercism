// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", "")

	answer := make([]string, 0)
	parts := strings.Split(s, " ")
	for _, c := range parts {
		if c != "" {
			answer = append(answer, strings.ToUpper(c[0:1]))
		}
	}
	return strings.Join(answer, "")
}
