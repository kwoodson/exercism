package sublist

type Relation string

const Equal Relation = "equal"
const Unequal Relation = "unequal"
const SubList Relation = "sublist"
const SuperList Relation = "superlist"

func compareList(a, b []int) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}
		return false
	}
	return true
}

func contains(a, b []int) bool {
	if len(a) == 0 {
		return true
	}
	for i := 0; i < len(b)-1; i++ {
		if i+len(a) > len(b) {
			return false
		}
		if a[0] == b[i] && compareList(a, b[i:i+len(a)]) {
			return true
		}
	}
	return false
}

func Sublist(first, second []int) Relation {
	// equal length, check for equality
	if len(first) == len(second) {
		if compareList(first, second) {
			return Equal
		}
		// check A is superlist of b
	} else if len(first) > len(second) {
		if contains(second, first) {
			return SuperList
		}
	} else if len(second) > len(first) {
		if contains(first, second) {
			return SubList
		}
	}
	return Unequal
}
