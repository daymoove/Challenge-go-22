package piscine

func IsUpper(s string) bool {
	l := []rune(s)
	for i := 0; i < len(s); i++ {
		if l[i] >= 'A' && l[i] <= 'Z' {
		} else {
			return false
		}
	}
	return true
}
