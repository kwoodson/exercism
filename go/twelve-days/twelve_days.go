package twelve

import (
	"fmt"
	"strings"
)

var DayMap = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

var gifts = map[int]string{
	1:  "a Partridge in a Pear Tree.",
	2:  "Turtle Doves",
	3:  "French Hens",
	4:  "Calling Birds",
	5:  "Gold Rings",
	6:  "Geese-a-Laying",
	7:  "Swans-a-Swimming",
	8:  "Maids-a-Milking",
	9:  "Ladies Dancing",
	10: "Lords-a-Leaping",
	11: "Pipers Piping",
	12: "Drummers Drumming",
}

var numMap = map[int]string{
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
}

func Verse(i int) string {
	sb := strings.Builder{}
	st := "On the %v day of Christmas my true love gave to me: "

	if i == 1 {
		sb.WriteString(fmt.Sprintf(st, DayMap[i]))
		sb.WriteString(gifts[i])
	} else {
		sb.WriteString(fmt.Sprintf(st, DayMap[i]))
		for z := i; z > 0; z-- {
			if z == 1 {
				sb.WriteString("and " + gifts[z])
			} else {
				sb.WriteString(fmt.Sprintf("%s %s, ", numMap[z], gifts[z]))
			}
		}
	}
	return sb.String()
}

func Song() string {
	sb := strings.Builder{}
	for i := 1; i < len(DayMap)+1; i++ {
		sb.WriteString(Verse(i))
		if i != len(DayMap) {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
