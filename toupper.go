package piscine

func ToUpper(s string) string {
	l := []rune(s)
	for i := 0; i < len(s); i++ {
		if l[i] >= 'a' && l[i] <= 'z' {
			l[i] = l[i] - 32
		}
	}
	z := string(l)
	return z
}
