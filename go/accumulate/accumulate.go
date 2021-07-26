package accumulate

func Accumulate(input []string, modifier func(string) string) []string {
	results := make([]string, 0)
	for _, s := range input {
		results = append(results, modifier(s))
	}
	return results
}
