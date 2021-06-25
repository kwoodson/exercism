// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	question, _ := regexp.Compile(`^.*\?$`)
	normal, _ := regexp.Compile(`[A-Za-z]+`)
	nowords, _ := regexp.Compile(`^\s+$`)
	remark = strings.Trim(remark, " ")
	rstr := ""
	if question.MatchString(remark) {
		if remark == strings.ToUpper(remark) && normal.MatchString(remark) {
			rstr = "Calm down, I know what I'm doing!"
		} else {
			rstr = "Sure."
		}
	} else if remark == strings.ToUpper(remark) && normal.MatchString(remark) {
		rstr = "Whoa, chill out!"
	} else if remark == "" || nowords.MatchString(remark) {
		rstr = "Fine. Be that way!"
	} else {
		rstr = "Whatever."
	}

	return rstr
}
