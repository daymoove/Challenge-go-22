package piscine

func ToLower(s string) string {
	l := []rune(s)
	for i := 0; i < len(s); i++ {
		if l[i] >= 'A' && l[i] <= 'Z' {
			l[i] = l[i] + 32
		}
	}
	z := string(l)
	return z
}
