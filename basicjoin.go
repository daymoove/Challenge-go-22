package piscine

func BasicJoin(elems []string) string {
	var m string
	for i := 0; i < len(elems); i++ {
		m = m + elems[i]
	}
	return m
}
