package series

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func All(n int, s string) []string {
	res := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if i+n > len(s) {
			break
		}
		res = append(res, UnsafeFirst(n, s[i:]))
	}
	return res
}
