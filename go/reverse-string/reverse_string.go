package reverse

func Reverse(s string) string {

	sb := ""
	for _, t := range s {
		sb = string(t) + sb
	}

	return sb
}
