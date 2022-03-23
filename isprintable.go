package piscine

func IsPrintable(s string) bool {
	l := []rune(s)
	for i := 0; i < len(s); i++ {
		if l[i] < 32 {
			return false
		}
	}
	return true
}
