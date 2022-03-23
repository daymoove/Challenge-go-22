package piscine

func AlphaCount(s string) int {
	l := []rune(s)
	m := 0
	for i := 0; i < len(s); i++ {
		if l[i] >= 'A' && l[i] <= 'Z' || l[i] >= 'a' && l[i] <= 'z' {
			m++
		}
	}
	return m
}
