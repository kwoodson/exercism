package house

var action = []string{
	"that lay in ",
	"that ate ",
	"that killed ",
	"that worried ",
	"that tossed ",
	"that milked ",
	"that kissed ",
	"that married ",
	"that woke ",
	"that kept ",
	"that belonged to ",
}

var noun = []string{
	"the house that Jack built.",
	"the malt",
	"the rat",
	"the cat",
	"the dog",
	"the cow with the crumpled horn",
	"the maiden all forlorn",
	"the man all tattered and torn",
	"the priest all shaven and shorn",
	"the rooster that crowed in the morn",
	"the farmer sowing his corn",
	"the horse and the hound and the horn",
}

func line(v int) string {
	if v-1 == 0 {
		return action[v-1] + noun[v-1]
	}

	return action[v-1] + noun[v-1] + "\n" + line(v-1)
}

func Verse(v int) string {
	start := "This is "
	if v == 1 {
		return start + noun[0]
	}

	return start + noun[v-1] + "\n" + line(v-1)
}

func Song() string {
	s := ""
	for i := 0; i < len(noun); i++ {
		s += Verse(i + 1)
		if i == len(noun)-1 {
			break
		}
		s += "\n\n"
	}
	return s
}
