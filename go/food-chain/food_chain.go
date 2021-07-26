package foodchain

import "fmt"

var animals = []string{
	"fly",
	"spider",
	"bird",
	"cat",
	"dog",
	"goat",
	"cow",
	"horse",
}

var aux = []string{
	"", // nothing on the spider
	"It wriggled and jiggled and tickled inside her.",
	"How absurd to swallow a bird!",
	"Imagine that, to swallow a cat!",
	"What a hog, to swallow a dog!",
	"Just opened her throat and swallowed a goat!",
	"I don't know how she swallowed a cow!",
}

var lines = []string{
	"I don't know why she swallowed the fly. Perhaps she'll die.",
	"that wriggled and jiggled and tickled inside her",
}

func line(n int) string {
	if n <= 1 {
		return lines[0]
	}
	z := fmt.Sprintf("She swallowed the %v to catch the %v", animals[n-1], animals[n-2])

	if animals[n-2] == "spider" {
		z += fmt.Sprintf(" %v", lines[1])
	}
	z += ".\n"
	return z + line(n-1)
}

func Verse(v int) string {
	start := "I know an old lady who swallowed a "
	if v == 1 {
		return fmt.Sprintf("%v%v.\n", start, animals[v-1]) + line(v-1)
	} else if v == 8 {
		return fmt.Sprintf("%v%v.\nShe's dead, of course!", start, animals[v-1])
	}

	return fmt.Sprintf("%v%v.\n%v\n", start, animals[v-1], aux[v-1]) + line(v)
}

func Verses(s, f int) string {
	res := ""
	for i := s; i <= f; i++ {
		if i == f {
			res += Verse(i)
			break
		}
		res += Verse(i) + "\n\n"
	}

	return res
}

func Song() string {
	return Verses(1, 8)
}
