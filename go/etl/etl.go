package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	out := map[string]int{}
	for iv, item := range in {
		for _, s := range item {
			out[strings.ToLower(s)] = iv
		}
	}
	return out
}
