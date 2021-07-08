package leap

// IsLeapYear returns whether year is leap year
func IsLeapYear(yr int) bool {
	if yr%4 == 0 {
		if yr%100 == 0 {
			return yr%400 == 0
		}
		return true
	}
	return false
}
