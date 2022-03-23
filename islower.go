package piscine

func IsLower(s string) bool {
	l := []rune(s)
	for i := 0; i < len(s); i++ {
		if l[i] >= 'a' && l[i] <= 'z' {
		} else {
			return false
		}
	}
	return true
}
