package robotname

import (
	"errors"
	"strings"
)

type Robot struct {
	name string
}

var names [26 * 26 * 10 * 10 * 10]string
var idx int

var letters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var numbers string = "0123456789"

//init build all string possibilities and place them in an array
func init() {
	sb := strings.Builder{}
	index := 0
	for a := 0; a < len(letters); a++ {
		for b := 0; b < len(letters); b++ {
			for c := 0; c < len(numbers); c++ {
				for d := 0; d < len(numbers); d++ {
					for e := 0; e < len(numbers); e++ {
						sb.WriteByte(letters[a])
						sb.WriteByte(letters[b])
						sb.WriteByte(numbers[c])
						sb.WriteByte(numbers[d])
						sb.WriteByte(numbers[e])
						names[index] = sb.String()
						index++
						sb.Reset()
					}
				}
			}
		}
	}
}

// Name fetch a name out of the names array and increment the
// global index counter. return error when names have been exhausted
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if idx == len(names) {
			return "", errors.New("Exhausted names")
		}
		r.name = names[idx]
		idx++
	}

	return r.name, nil
}

// Reset set name to empty string
func (r *Robot) Reset() {
	r.name = ""
}

// Original Name() function that called random name fetching and
// kept track of used ones in a map. Very inefficient as can be
// seen with the collision messages.
// func (r *Robot) Name() (string, error) {
// 	if r.name == "" {
// 		for {
// 			n, err := generateName()
// 			if err != nil {
// 				return "", nil
// 			}
// 			r.name = n
// 			if !used[r.name] {
// 				break
// 			}
// 			//  else {
// 			// fmt.Println("Collision")
// 			// }
// 		}
// 	}
// 	if len(used) == totalNames {
// 		return "", errors.New("Exhausted the names")
// 	}

// 	used[r.name] = true
// 	return r.name, nil
// }

//generateName selects random numbers to generate a name
// first naive approach was successful but inefficient
// func generateName() (string, error) {
// 	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	numbers := "0123456789"

// 	r := make([]byte, 0)
// 	for i := 0; i < 2; i++ {
// 		o, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
// 		if err != nil {
// 			return "", err
// 		}
// 		r = append(r, letters[o.Int64()])
// 	}
// 	for i := 0; i < 3; i++ {
// 		o, err := rand.Int(rand.Reader, big.NewInt(int64(len(numbers))))
// 		if err != nil {
// 			return "", err
// 		}
// 		r = append(r, numbers[o.Int64()])
// 	}
// 	return string(r), nil
// }
