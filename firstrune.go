package piscine

func FirstRune(s string) rune {
	var m rune
	for _, letter := range s {
		m = letter
		break
	}
	return m
}
