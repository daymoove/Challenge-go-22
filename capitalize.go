package piscine

func Capitalize(s string) string {
	l := []rune(s)
	if l[0] >= 'a' && l[0] <= 'z' {
		l[0] = l[0] - 32
	}
	for i := 1; i < len(s); i++ {
		if l[i-1] >= 'a' && l[i-1] <= 'z' || l[i-1] >= 'A' && l[i-1] <= 'Z' || l[i-1] >= '0' && l[i-1] <= '9' {
			if l[i] >= 'A' && l[i] <= 'Z' {
				l[i] = l[i] + 32
			}
		} else if l[i] >= 'a' && l[i] <= 'z' {
			l[i] = l[i] - 32
		}
	}
	m := string(l)
	return m
}
