package internal

// AbsInt takes an integer and returns it's absolute value.
// For example: -4 returns 4, 5 returns 5, etc.
func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
