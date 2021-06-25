package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(data []string) FreqMap {
	rc := make(chan FreqMap)

	for _, s := range data {
		go func(s string) {
			rc <- Frequency(s)
		}(s)
	}

	results := FreqMap{}
	for range data {
		for k, v := range <-rc {
			results[k] += v
		}
	}
	return results
}
