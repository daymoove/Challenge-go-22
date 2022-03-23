package piscine

func IsNumeric(s string) bool {
	l := []rune(s)
	for i := 0; i < len(s); i++ {
		if l[i] >= '0' && l[i] <= '9' {
		} else {
			return false
		}
	}
	return true
}
