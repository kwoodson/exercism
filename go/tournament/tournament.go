package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type Team struct {
	win    int
	loss   int
	draw   int
	name   string
	points int
	played int
}

func validOutcome(o string) bool {
	if o == "win" || o == "loss" || o == "draw" {
		return true
	}
	return false
}

func Add(t Team, outcome string) Team {
	switch outcome {
	case "draw":
		t.draw++
		t.points++
	case "win":
		t.win++
		t.points += 3
	case "loss":
		t.loss++
	}
	t.played++
	return t
}

func Tally(r io.Reader, buffer io.Writer) error {
	var teams = map[string]Team{}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	for _, s := range strings.Split(string(data), "\n") {
		if s == "" || strings.HasPrefix(s, "#") {
			continue
		}
		parts := strings.Split(s, ";")
		if len(parts) != 3 {
			return errors.New("Malformed results line")
		}
		t1 := parts[0] // team1 name
		t2 := parts[1] // team2 name
		o := parts[2]  // outcome
		if !validOutcome(o) {
			return errors.New("Error in the outcome")
		}

		if _, ok := teams[t1]; !ok {
			teams[t1] = Team{
				name: t1,
			}
		}
		o2 := o
		if o == "win" {
			o2 = "loss"
		} else if o == "loss" {
			o2 = "win"
		}
		if _, ok := teams[t2]; !ok {
			teams[t2] = Team{
				name: t2,
			}
		}
		teams[t1] = Add(teams[t1], o)
		teams[t2] = Add(teams[t2], o2)
	}
	io.Copy(buffer, strings.NewReader(buildResults(teams)))
	return nil
}

func buildResults(teams map[string]Team) string {
	sb := strings.Builder{}
	sb.WriteString("Team                           | MP |  W |  D |  L |  P\n")

	// make a struct to sort map[string]Team by
	// name and total points.
	type kv struct {
		Key   string
		Value int
	}
	t := make([]kv, 0)
	for x := range teams {
		tmp := teams[x]
		t = append(t, kv{x, tmp.points})
	}

	sort.Slice(t, func(i, j int) bool {
		if t[i].Value == t[j].Value { // if tied, sort alphabetically
			return t[i].Key < t[j].Key
		}
		return t[i].Value > t[j].Value // reverse sort by total points
	})

	for _, i := range t {
		tmp := teams[i.Key]
		sb.WriteString(fmt.Sprintf("%-31v|  %v |  %v |  %v |  %v |  %v\n",
			tmp.name,
			tmp.played,
			tmp.win,
			tmp.draw,
			tmp.loss,
			tmp.points,
		))
	}
	return sb.String()
}
